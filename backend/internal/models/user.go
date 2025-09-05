package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"password_hash"`
	Role         string    `gorm:"not null" json:"role"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
