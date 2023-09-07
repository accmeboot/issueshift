package respository

import (
	"context"
	"github.com/accmeboot/issueshift/internal/domain"
	"time"
)

func (p *Provider) CreateTask(title, description, status string, assignee int64) error {
	query := `INSERT INTO tasks (title, description, status, assignee) VALUES ($1, $2, $3, $4)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := p.db.ExecContext(ctx, query, title, description, status, assignee)
	if err != nil {
		return domain.ErrServer(err)
	}

	return nil
}

func (p *Provider) GetAllTasks() ([]*domain.Task, error) {
	query := `
			  SELECT tasks.title, tasks.description, tasks.created_at, tasks.updated_at, tasks.status, users.email
			  FROM tasks INNER JOIN users
			  ON tasks.assignee = users.id;
			 `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, domain.ErrServer(err)
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	var tasks []*domain.Task

	for rows.Next() {
		var task domain.Task
		var email string

		err := rows.Scan(
			&task.Title,
			&task.Description,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.Status,
			&email,
		)

		if err != nil {
			return nil, domain.ErrServer(err)
		}

		task.Assignee = &domain.User{
			Email: email,
		}

		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		return nil, domain.ErrServer(err)
	}

	return tasks, nil
}
