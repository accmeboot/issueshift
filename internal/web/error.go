package web

import (
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

func (h *Handler) ErrorView(w http.ResponseWriter, r *http.Request) {
	h.PagesCache.Render(
		w,
		http.StatusInternalServerError,
		"error.gohtml",
		nil,
		domain.Envelope{"error": r.Context().Value(domain.ErrKey)},
	)
}
