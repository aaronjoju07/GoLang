package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"password" validate:"main=6,required"`
	Updated_at    time.Time          `json:"updated_at"`
	Email         *string            `json:"email" validate:"email,required"`
	Created_at    time.Time          `json:"created_at"`
	Avatar        *string            `json:"avatar"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token" `
	Refresh_Token *string            `json:"refresh_token" `
	User_id       string             `json:"user_id" `
}
