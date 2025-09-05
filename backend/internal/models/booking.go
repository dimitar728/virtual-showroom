package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	ShowroomID uuid.UUID `gorm:"type:uuid;not null" json:"showroom_id"`
	SlotTime   time.Time `json:"slot_time"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
