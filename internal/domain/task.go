package domain

import "time"

type Task struct {
	ID          int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Assignee    *User
}

type TaskService interface {
	Create(title, description string) error
}

type TaskRepository interface {
	Create(title, description string) error
}
