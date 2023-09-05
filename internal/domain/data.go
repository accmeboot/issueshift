package domain

import (
	"time"
)

type userKey string
type errKey string

const UserKey = userKey("UserKey")
const ErrKey = errKey("appError")

type Envelope map[string]any

type ErrNoRecord error
type ErrAlreadyExists error
type ErrEditConflict error
type ErrServer error
type ErrInvalidCredentials error
type ErrBadlyFormattedJson error

type Image struct {
	ID        int64
	Name      string
	ImageData []byte
}

type Task struct {
	ID          int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Assignee    *User
	Status      string
}

type Token struct {
	plain  *string
	hash   []byte
	UserID int64
	Expiry time.Time
}

type User struct {
	ID           int64
	Email        string
	PasswordHash []byte
	Name         string
	CreatedAt    time.Time
	AvatarID     *int64
}
