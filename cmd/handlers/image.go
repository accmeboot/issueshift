package handlers

import (
	"errors"
	"fmt"
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
		p.helpers.SendUnprocessableEntity(w, err)
		return
	}

	image, err := p.service.GetImage(id)
	if err != nil {
		var noRecord *domain.ErrNoRecord
		switch {
		case errors.As(err, noRecord):
			p.helpers.SendNotFound(w, err)
		default:
			p.helpers.SendServerError(w, err)
		}
		return
	}
	w.Header().Set("Content-Type", "image/png")
	if _, err = w.Write(image.ImageData); err != nil {
		log.Println(err)
	}
}

func (p *Provider) CreateImage(w http.ResponseWriter, r *http.Request) {
	fileType := r.Header.Get("Content-Type")

	if fileType != "image/png" && fileType != "image/jpeg" && fileType != "image.ts/svg" {
		p.helpers.SendBadRequest(w, domain.Error{
			"image": fmt.Sprintf("filetype: %s is not allowed", fileType),
		}, nil)
		return
	}

	id, err := p.service.CreateImage(r.Body)
	if err != nil {
		p.helpers.SendServerError(w, err)
		return
	}

	p.helpers.Send(w, http.StatusCreated, domain.Envelope{"id": id})
}
