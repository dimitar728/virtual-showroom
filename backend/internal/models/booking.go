package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookingStatus string

const (
	StatusPending   BookingStatus = "pending"
	StatusConfirmed BookingStatus = "confirmed"
	StatusCancelled BookingStatus = "cancelled"
)

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	b.CreatedAt = time.Now()
type Booking struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	ShowroomID uuid.UUID `gorm:"type:uuid;not null" json:"showroom_id"`
	SlotTime   time.Time `gorm:"not null" json:"slot_time"`
	Status     string    `gorm:"default:'pending'" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	b.CreatedAt = time.Now()

	return
}
