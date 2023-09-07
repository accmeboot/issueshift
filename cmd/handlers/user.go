package handlers

import (
	"errors"
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

func (p *Provider) WhoAmI(w http.ResponseWriter, r *http.Request) {
	rawUser := r.Context().Value(domain.UserKey)

	if rawUser != nil {
		user := rawUser.(*domain.User)

		p.helpers.Send(w, http.StatusOK, domain.Envelope{
			"user": user,
		})

		return
	}

	p.helpers.SendServerError(w, errors.New("there is no user in the context"))
}
