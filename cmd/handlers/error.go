package handlers

import (
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

func (p *Provider) ErrorView(w http.ResponseWriter, r *http.Request) {
	p.pages.Render(
		w,
		http.StatusInternalServerError,
		"error.gohtml",
		nil,
		domain.Envelope{"error": r.Context().Value(domain.ErrKey)},
	)
}
