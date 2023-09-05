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

func (p *Provider) GetAllTasks() ([]*domain.Task, error) {
	query := `
			  SELECT
					tasks.title,
			        CASE 
						 WHEN LENGTH(tasks.description) > 150 THEN LEFT(tasks.description, 150) || '...' 
						 ELSE tasks.description 
					END AS description_truncated,
			    tasks.created_at, tasks.updated_at, tasks.status, users.email
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
