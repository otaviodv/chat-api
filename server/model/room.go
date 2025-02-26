package model

import (
	"chat-api/db"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var roomCollection *mongo.Collection

func init() {
	roomCollection = db.GetCollection("rooms")
}

type Room struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"title" bson:"title"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

type RoomItem struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"title" bson:"title"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	Messages  []Message `json:"messages" bson:"messages"`
}

func (r *Room) Create() (*primitive.ObjectID, error) {
	r.Id = primitive.NewObjectID()
	r.CreatedAt = time.Now()
	result, err := roomCollection.InsertOne(context.Background(), r)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)
	return &id, nil
}

func (r *RoomItem) Get(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	q := bson.M{"_id": objID}
	err = roomCollection.FindOne(context.Background(), q).Decode(r)
	if err != nil {
		return err
	}

	r.Messages, err = getLastMessagesFromRoom(r.Id)

	return err

}
