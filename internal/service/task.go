package service

func (p *Provider) CreateTask(title, description string) error {
	return p.repository.CreateTask(title, description)
}
