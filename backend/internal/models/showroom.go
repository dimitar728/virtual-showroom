package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Showroom struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	ModelPath   string    `gorm:"not null" json:"model_path"`
	Capacity    int       `json:"capacity"`
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// Auto-generate UUID before insert
func (s *Showroom) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
