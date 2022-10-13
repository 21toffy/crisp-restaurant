package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required, min=3, max=100"`
	Category   string             `json:"category" validate:"required"`
	Start_Date time.Time          `json:"updated_at"`
	End_Date   time.Time          `json:"updated_at"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Menu_id    string             `json:"food_id"`
}

// Menu_id    string             `json:"menu_id"`
