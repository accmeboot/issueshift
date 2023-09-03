package respository

import (
	"context"
	"database/sql"
	"github.com/accmeboot/issueshift/internal/domain"
	"time"
)

type TaskRepository struct {
	DB *sql.DB
}

var _ domain.TaskRepository = &TaskRepository{}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (tr *TaskRepository) Create(title, description string) error {
	query := `INSERT INTO tasks (title, description) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := tr.DB.ExecContext(ctx, query, title, description)
	if err != nil {
		return domain.ErrServer(err)
	}

	return nil
}
