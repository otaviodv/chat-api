package db

import (
	"context"
	"fmt"
	"log"
	"os"
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

	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")
	addr := os.Getenv("MONGO_ADDR")

	uri := fmt.Sprintf("mongodb://%s:%s@%s/?authSource=admin&readPreference=primary&directConnection=true&SSL=false", user, pass, addr)
	connect(ctx, uri, "chat-api")
}

func connect(ctx context.Context, uri string, database string) {
	var err error

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Println("Mongo Connection error ==>", err)
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("Mongo PING error==>", err)
		log.Fatal(err)
	}
	log.Println("Database connection succeeded")
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
