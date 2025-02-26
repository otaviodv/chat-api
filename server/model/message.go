package model

import (
	"chat-api/db"
	"context"
	"slices"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var messageCollection *mongo.Collection

func init() {
	messageCollection = db.GetCollection("messages")
}

type Message struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	RoomId   primitive.ObjectID `json:"roomId" bson:"roomId"`
	Username string             `json:"username" bson:"username"`
	Text     string             `json:"text" bson:"text"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

func (m *Message) Create() (*primitive.ObjectID, error) {
	m.Id = primitive.NewObjectID()
	m.CreatedAt = time.Now()
	result, err := messageCollection.InsertOne(context.Background(), m)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)
	return &id, nil
}

func getLastMessagesFromRoom(roomId primitive.ObjectID) ([]Message, error) {
	filter := bson.M{"roomId": roomId}
	options := options.Find().SetLimit(50)
	options.SetSort(bson.D{{Key: "_id", Value: -1}})

	cursor, err := messageCollection.Find(context.Background(), filter, options)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []Message{}, nil
		}
		return nil, err
	}
	var messages []Message

	if err = cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}
	slices.Reverse(messages)

	return messages, nil
}
