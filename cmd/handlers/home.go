package handlers

import (
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

func (p *Provider) HomeView(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(domain.UserKey).(*domain.User)

	allTasks, err := p.service.GetAllTasks()
	if err != nil {
		p.pages.ServerError(w, err)
		return
	}

	var todoTasks []*domain.Task
	var inProgressTasks []*domain.Task
	var doneTasks []*domain.Task

	for _, task := range allTasks {
		switch task.Status {
		case "todo":
			todoTasks = append(todoTasks, task)
		case "in_progress":
			inProgressTasks = append(inProgressTasks, task)
		case "done":
			doneTasks = append(doneTasks, task)
		default:
			break
		}
	}

	p.pages.Render(w, http.StatusOK, "home.gohtml", nil, domain.Envelope{
		"IsAuthenticated": ok,
		"Todo":            todoTasks,
		"InProgress":      inProgressTasks,
		"Done":            doneTasks,
	})
}
