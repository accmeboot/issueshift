package web

import (
	"context"
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

func (h *Handler) Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("Bearer")
		if err != nil {
			w.Header().Set("HX-Redirect", "/signin")
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}

		user, err := h.TokenService.Authenticate(tokenCookie.Value)
		if err != nil {
			w.Header().Set("HX-Redirect", "/signin")
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}

		*r = *r.WithContext(context.WithValue(r.Context(), domain.UserKey, user))

		next.ServeHTTP(w, r)
	})
}
