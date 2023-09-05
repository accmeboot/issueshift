package handlers

import (
	"errors"
	"github.com/accmeboot/issueshift/internal/domain"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func (p *Provider) GetImage(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(rawId, 16, 64)
	if err != nil {
		//response.WriteError(w, http.StatusUnprocessableEntity, domain.Envelope{"error": "id is not valid"}, err)
		return
	}

	image, err := p.service.GetImage(id)
	if err != nil {
		var noRecord domain.ErrNoRecord
		switch {
		case errors.As(err, &noRecord):
			w.WriteHeader(http.StatusNotFound)
		default:
			p.pages.ServerError(w, err)
		}
		return
	}
	w.Header().Set("Content-Type", "image/png")
	if _, err = w.Write(image.ImageData); err != nil {
		log.Println(err)
	}
}
