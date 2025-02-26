package controller

import "go.mongodb.org/mongo-driver/bson/primitive"

type createResponse struct {
	Id primitive.ObjectID `json:"id,omitempty"`
}
