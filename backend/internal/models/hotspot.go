package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hotspot struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ShowroomID  uuid.UUID `gorm:"type:uuid;not null;index" json:"showroom_id"`
	Label       string    `gorm:"not null" json:"label"`
	Description string    `json:"description"`
	// World-space position for the marker in the model's coordinate system
	PosX float64 `gorm:"not null" json:"pos_x"`
	PosY float64 `gorm:"not null" json:"pos_y"`
	PosZ float64 `gorm:"not null" json:"pos_z"`
	// Optional extras
	MediaURL  string    `json:"media_url"`
	LinkURL   string    `json:"link_url"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (h *Hotspot) BeforeCreate(tx *gorm.DB) (err error) {
	h.ID = uuid.New()
	return
}
