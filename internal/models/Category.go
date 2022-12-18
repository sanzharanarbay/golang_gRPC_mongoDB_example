package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Label     string             `json:"label" binding:"required" bson:"label,omitempty"`
	Total     int64              `json:"total" binding:"required"  bson:"total,omitempty"`
	Status    bool               `json:"status" binding:"required" bson:"status,omitempty"`
	Order     int64              `json:"order" binding:"required" bson:"order,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
