package domain

import "mime/multipart"

type ServiceProvider interface {
	CreateImage(file *multipart.File, fileName string) (int64, error)
	GetImage(id int64) (*Image, error)
	CreateTask(title, description, status string, assignee int64) error
	GetAllTasks() ([]*Task, error)
	Authenticate(token string) (*User, error)
	CreateToken(userId int64) (*string, error)
	GetUser(email, password string) (*User, error)
	CreateUser(email, name, password string, avatarId *int64) error
}
