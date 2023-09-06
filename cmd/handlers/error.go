package handlers

import (
	"github.com/accmeboot/issueshift/cmd/helpers"
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

func (p *Provider) ErrorView(w http.ResponseWriter, r *http.Request) {
	p.templates.Render(helpers.RenderDTO{
		Writer:   w,
		Status:   http.StatusInternalServerError,
		Template: "error.gohtml",
		Data: domain.Envelope{
			"error": r.Context().Value(domain.ErrKey),
		},
	})
}
