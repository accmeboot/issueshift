package service

import "github.com/accmeboot/issueshift/internal/domain"

func (p *Provider) CreateTask(title, description string) error {
	return p.repository.CreateTask(title, description)
}

func (p *Provider) GetAllTasks() ([]*domain.Task, error) {
	return p.repository.GetAllTasks()
}
