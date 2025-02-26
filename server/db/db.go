package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	dbname string
)

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	connect(ctx, "mongodb://localhost:27017", "chat-api")
}

func connect(ctx context.Context, uri string, database string) {
	var err error

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectando Mongo")
	// err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Erro connect==>", err)
		log.Fatal(err)
	}

	fmt.Println("Ping")
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Erro ping==>", err)
		log.Fatal(err)
	}
	dbname = database
}

func GetCollection(name string) *mongo.Collection {
	return client.Database(dbname).Collection(name)
}

func Disconnect(ctx context.Context) {
	client.Disconnect(ctx)
}

func DropDatabase(ctx context.Context) error {
	err := client.Database(dbname).Drop(ctx)
	return err
}
