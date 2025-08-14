package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	BookingStatusPending   = "pending"
	BookingStatusConfirmed = "confirmed"
	BookingStatusCancelled = "cancelled"
)

type Booking struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;index:idx_bookings_slot" json:"user_id"`
	ShowroomID uuid.UUID `gorm:"type:uuid;not null;index:idx_bookings_slot" json:"showroom_id"`
	SlotTime   time.Time `gorm:"not null;index:idx_bookings_slot" json:"slot_time"`
	Status     string    `gorm:"type:varchar(20);default:'pending';index" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}
