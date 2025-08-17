package models

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"

	StatusActive    UserStatus = "active"
	StatusSuspended UserStatus = "suspended"
)
