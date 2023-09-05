package respository

import (
	"context"
	"github.com/accmeboot/issueshift/internal/domain"
	"time"
)

func (p *Provider) CreateTask(title, description string) error {
	query := `INSERT INTO tasks (title, description) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := p.db.ExecContext(ctx, query, title, description)
	if err != nil {
		return domain.ErrServer(err)
	}

	return nil
}
