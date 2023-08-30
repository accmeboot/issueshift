package domain

import (
	"time"
)

type Token struct {
	plain  string
	UserID int64
	Expiry time.Time
}

type TokenRepository interface {
	GetUserFromToken(hash []byte) (*User, error)
	Create(userId int64, hash []byte, expiry time.Time) error
}

type TokenService interface {
	Authenticate(token string) (*User, error)
	Create(userId int64) (*string, error)
}
