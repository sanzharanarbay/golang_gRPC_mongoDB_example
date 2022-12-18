package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Vacancy struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title" binding:"required" bson:"title,omitempty"`
	Slug         string             `json:"slug" binding:"required"  bson:"slug,omitempty"`
	Description  string             `json:"description" binding:"required" bson:"description,omitempty"`
	Salary       string             `json:"salary" binding:"required"  bson:"salary,omitempty"`
	IsActive     bool               `json:"is_active" binding:"required" bson:"is_active,omitempty"`
	IsDistribute bool               `json:"is_distribute" binding:"required" bson:"is_distribute,omitempty"`
	Order        int64              `json:"order" binding:"required" bson:"order,omitempty"`
	Cities       []*City            `json:"cities" bson:"cities,omitempty"`
	Categories   []*Category        `json:"categories" bson:"categories,omitempty"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type VacancyRequestBody struct {
	Title        string   `json:"title" binding:"required"`
	Description  string   `json:"description" binding:"required"`
	Salary       string   `json:"salary,omitempty"`
	IsActive     bool     `json:"is_active,omitempty"`
	IsDistribute bool     `json:"is_distribute,omitempty"`
	Order        int64    `json:"order,omitempty"`
	Cities       []string `json:"cities" binding:"required"`
	Categories   []string `json:"categories" binding:"required"`
}
