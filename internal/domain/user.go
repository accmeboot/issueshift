package domain

import "github.com/accmeboot/issueshift/internal/_model"

type User struct {
	_model.User
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	CreateUser(email, name string, avatarUrl *string, passwordHash []byte) error
}
