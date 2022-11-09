package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `json:"title" `
	Text  string             `json:"text" `
}
