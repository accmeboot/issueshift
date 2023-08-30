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
	AvatarID     *int64
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	CreateUser(email, name string, avatarId *int64, passwordHash []byte) error
}

type UserService interface {
	GetUserByCredentials(email, password string) (*User, error)
	CreateUser(email, name, password string, avatarId *int64) error
}
