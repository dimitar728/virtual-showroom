package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"type:varchar(20);default:'user'" json:"role"`
	Suspended    bool      `gorm:"default:false" json:"suspended"`
	CreatedAt    time.Time `json:"created_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

type Showroom struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	ModelPath   string    `json:"model_path"`
	Capacity    int       `json:"capacity"`
	CreatedBy   uuid.UUID `gorm:"type:uuid" json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

func (s *Showroom) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}

type Booking struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	ShowroomID uuid.UUID `gorm:"type:uuid;not null" json:"showroom_id"`
	SlotTime   time.Time `json:"slot_time"`
	Status     string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}
