package domain

import (
	"time"
)

type User struct {
	ID           int64
	Email        string
	PasswordHash []byte
	Name         string
	CreatedAt    time.Time
	AvatarUrl    *string
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	CreateUser(email, name string, avatarUrl *string, passwordHash []byte) error
}
