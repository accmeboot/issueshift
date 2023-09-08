package handlers

import (
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

type CreateTaskDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Assignee    int64  `json:"assignee" validate:"required"`
	Status      string `json:"status" validate:"task_status"`
}

func (p *Provider) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	tasks, err := p.service.GetAllTasks(status)
	if err != nil {
		p.helpers.SendServerError(w, err)
		return
	}

	p.helpers.Send(w, http.StatusOK, domain.Envelope{
		"tasks": tasks,
	})
}

func (p *Provider) CreateTask(w http.ResponseWriter, r *http.Request) {
	var DTO CreateTaskDTO

	err := p.helpers.ReadBody(w, r, &DTO)
	if err != nil {
		p.helpers.SendUnprocessableEntity(w, err)
		return
	}

	validator := p.helpers.NewValidator()

	if len(DTO.Status) < 1 {
		DTO.Status = "todo"
	}

	if ok := validator.Validate(DTO); !ok {
		p.helpers.SendBadRequest(w, validator.Errors, nil)
		return
	}

	err = p.service.CreateTask(DTO.Title, DTO.Description, DTO.Status, DTO.Assignee)
	if err != nil {
		p.helpers.SendServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
