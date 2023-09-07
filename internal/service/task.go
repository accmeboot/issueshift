package service

import "github.com/accmeboot/issueshift/internal/domain"

func (p *Provider) CreateTask(title, description, status string, assignee int64) error {
	return p.repository.CreateTask(title, description, status, assignee)
}

func (p *Provider) GetAllTasks() ([]*domain.Task, error) {
	return p.repository.GetAllTasks()
}
