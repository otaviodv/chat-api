package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommandTransport struct {
	RoomId   primitive.ObjectID `json:"roomId"`
	Username string             `json:"username"`
	Text     string             `json:"text"`
}
