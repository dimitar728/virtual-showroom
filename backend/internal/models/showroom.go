package models

import (
	"time"

	"github.com/google/uuid"
)

type Showroom struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ModelPath   string    `json:"model_path"`
	Capacity    int       `json:"capacity"`
	CreatedBy   uuid.UUID `gorm:"type:uuid" json:"created_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
