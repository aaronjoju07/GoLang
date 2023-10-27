package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_Date time.Time          `json:"order_date" validate:"required"`
	Order_id   string             `json:"order_id" `
	Updated_at time.Time          `json:"updated_at"`
	Created_at time.Time          `json:"created_at"`
	Table_id   *string            `json:"table_id" validate:"required"`
}
