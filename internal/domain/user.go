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
	Avatar       *string
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	CreateUser(email, name string, avatarUrl *string, passwordHash []byte) error
}

type UserService interface {
	GetUserByCredentials(email, password string) (*User, error)
	CreateUser(email, name string, avatarUrl *string, passwordHash []byte) error
}
