package middlewares

import (
	"context"
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
	"strings"
)

func (p *Provider) Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.Header.Get("Authorization"), " ")

		if len(parts) < 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		user, err := p.service.Authenticate(parts[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		*r = *r.WithContext(context.WithValue(r.Context(), domain.UserKey, user))

		next.ServeHTTP(w, r)
	})
}
