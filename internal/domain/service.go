package domain

import "io"

type ServiceProvider interface {
	CreateImage(file io.ReadCloser) (int64, error)
	GetImage(id int64) (*Image, error)
	CreateTask(title, description, status string, assignee int64) error
	GetAllTasks(status string) ([]*Task, error)
	Authenticate(token string) (*User, error)
	CreateToken(userId int64) (*string, error)
	GetUser(email, password string) (*User, error)
	CreateUser(email, name, password string, avatarId *int64) error
}
