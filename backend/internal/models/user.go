package models

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `json:"-"`
	Role         Role      `gorm:"type:varchar(10);not null;default:user" json:"role"`
	Suspended    bool      `json:"suspended"`
	CreatedAt    time.Time `json:"created_at"`
}

func (u *User) BeforeCreate(tx any) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
