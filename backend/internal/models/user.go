package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
<<<<<<< HEAD
	StatusActive    UserStatus = "active"
	StatusSuspended UserStatus = "suspended"
)
=======
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         UserRole  `gorm:"type:varchar(10);not null;default:'user'" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
=======
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"not null;default:user"`
	Status    string    `gorm:"not null;default:active"` // active | suspended | deleted
	CreatedAt time.Time
	UpdatedAt time.Time
>>>>>>> feature/HVSBS-110/UserManager-Suspend/delete/reactivate-users
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
<<<<<<< HEAD

=======
>>>>>>> feature/HVSBS-110/UserManager-Suspend/delete/reactivate-users
