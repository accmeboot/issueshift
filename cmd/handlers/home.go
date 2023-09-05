package handlers

import (
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

func (p *Provider) HomeView(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(domain.UserKey).(*domain.User)
	p.pages.Render(w, http.StatusOK, "home.gohtml", nil, domain.Envelope{"IsAuthenticated": ok})
}
