package models

import (
	"time"

	"github.com/google/uuid"
)

type Showroom struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	ModelPath   string    `json:"model_path"`
	Capacity    int       `json:"capacity"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

func (s *Showroom) BeforeCreate(tx any) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}
