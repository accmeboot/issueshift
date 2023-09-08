package domain

import (
	"time"
)

type userKey string

const UserKey = userKey("UserKey")

type Envelope map[string]any
type Error map[string]string

type ErrNoRecord error
type ErrAlreadyExists error
type ErrEditConflict error
type ErrServer error
type ErrInvalidCredentials error
type ErrBadlyFormattedJson error

type Image struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ImageData []byte `json:"image_data"`
}

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Assignee    *User     `json:"assignee"`
	Status      string    `json:"status"`
}

type Token struct {
	plain  *string
	hash   []byte
	UserID int64
	Expiry time.Time
}

type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	AvatarID     *int64    `json:"avatar_id"`
}
