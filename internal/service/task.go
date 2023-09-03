package service

import "github.com/accmeboot/issueshift/internal/domain"

type TaskService struct {
	repo domain.TaskRepository
}

var _ domain.TaskService = &TaskService{}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (ts *TaskService) Create(title, description string) error {
	return ts.repo.Create(title, description)
}
