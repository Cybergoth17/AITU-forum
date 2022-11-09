package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AdminPost struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `json:"title" `
	Text  string             `json:"text" `
}
