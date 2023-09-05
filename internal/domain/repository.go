package domain

import "time"

type RepositoryProvider interface {
	CreateImage(base64File, name string) (int64, error)
	GetImage(id int64) (*string, *string, error)
	CreateUser(email, name string, avatarId *int64, passwordHash []byte) error
	GetUser(email string) (*User, error)
	CreateToken(userId int64, hash []byte, expiry time.Time) error
	GetUserFromToken(token []byte) (*User, error)
	CreateTask(title, description string) error
}
