package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model that governs all notes objects retrived or inserted into the DB
type User struct {
	Visitor
	ID            primitive.ObjectID `bson:"_id"`
	Username      *string            `json:"username" bson:"username" form="name"`
	Password      *string            `json:"Password" validate:"required,min=6" bson:"password" form="password"`
	Email         *string            `json:"email" validate:"email,required" bson:"email" form="email"`
	Phone         *string            `json:"phone" form="phone"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
	IsSub         string             `json:"isSub"`
}
type UserInfo struct {
	Username string `json:"username" bson:"username" form="name"`
	Email    string `json:"email" validate:"email,required" bson:"email" form="email"`
	Phone    string `json:"phone" form="phone"`
	IsSub    bool   `json:"isSub"`
	IsLux    bool   `json:"isLux"`
}
