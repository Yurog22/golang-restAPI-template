package models

import (
	"github.com/google/uuid"
	"time"
)

type Example struct {
	ID          uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(); primary_key; unique" validate:"required" json:"id" example:"2badcae8-3b24-45f1-bdda-d8ed87c3d002"`
	StringValue string    `json:"stringValue" example:"string"`
	IntValue    int       `json:"intValue" example:"1"`
	CreatedAt   time.Time `gorm:"<-:create" json:"createdAt" example:"2024-06-19T17:05:21.418722+07:00"`
	UpdatedAt   time.Time `json:"updatedAt" example:"2024-06-19T17:05:21.418722+07:00"`
}
