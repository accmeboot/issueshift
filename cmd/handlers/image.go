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
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		p.helpers.SendUnprocessableEntity(w, err)
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		p.helpers.SendBadRequest(w, domain.Envelope{"error": "no files have been provided in the filed image"}, err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	id, err := p.service.CreateImage(&file, header.Filename)
	if err != nil {
		p.helpers.SendServerError(w, err)
		return
	}

	p.helpers.Send(w, http.StatusCreated, domain.Envelope{"image_id": id})
}
