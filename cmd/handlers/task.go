package handlers

import (
	"github.com/accmeboot/issueshift/cmd/helpers"
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

func (p *Provider) CreateTaskView(w http.ResponseWriter, r *http.Request) {
	p.templates.Render(helpers.RenderDTO{
		Writer:   w,
		Template: "create-task.gohtml",
		Name:     "create_task_popup",
		Data: domain.Envelope{
			"IsOpen": true,
		},
	})
}

func (p *Provider) CreatTaskViewCancel(w http.ResponseWriter, r *http.Request) {
	p.templates.Render(helpers.RenderDTO{
		Writer:   w,
		Template: "create-task.gohtml",
		Name:     "create_task_popup",
	})
}
