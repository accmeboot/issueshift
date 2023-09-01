package api

import (
	"errors"
	"fmt"
	"github.com/accmeboot/issueshift/internal/api/response"
	"github.com/accmeboot/issueshift/internal/domain"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type ImageHandler struct {
	imageService domain.ImageService
}

func NewImageHandler(is domain.ImageService) *ImageHandler {
	return &ImageHandler{imageService: is}
}

func (ih *ImageHandler) GetImage(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(rawId, 16, 64)
	if err != nil {
		response.WriteError(w, http.StatusUnprocessableEntity, domain.Envelope{"error": "id is not valid"}, err)
		return
	}

	image, err := ih.imageService.Get(id)
	if err != nil {
		var noRecord *domain.ErrNoRecord
		switch {
		case errors.As(err, noRecord):
			response.WriteError(
				w,
				http.StatusNotFound,
				domain.Envelope{"error": fmt.Sprintf("image with id: %d could not be found", id)},
				err,
			)
		default:
			response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "internal server error"}, err)
		}
		return
	}
	w.Header().Set("Content-Type", "image/png")
	if _, err = w.Write(image.ImageData); err != nil {
		log.Println(err)
	}
}

func (ih *ImageHandler) CreateImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "file is too big"}, err)
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, domain.Envelope{"error": "no files have been provided in the filed image"}, err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	id, err := ih.imageService.Create(&file, header.Filename)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "internal server error"}, err)
		return
	}

	response.WriteJSON(w, http.StatusCreated, domain.Envelope{"image_id": id})
}
