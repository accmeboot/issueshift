package web

import "net/http"

func (h *Handler) HomeView(w http.ResponseWriter, _ *http.Request) {
	h.PagesCache.Render(w, http.StatusOK, "home.gohtml", nil, nil)
}
