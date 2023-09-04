package web

import (
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

func (h *Handler) HomeView(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(domain.UserKey).(*domain.User)
	h.PagesCache.Render(w, http.StatusOK, "home.gohtml", nil, domain.Envelope{"IsAuthenticated": ok})
}
