package web

import (
	"github.com/accmeboot/issueshift/internal/domain"
	"github.com/accmeboot/issueshift/internal/web/request"
	"github.com/accmeboot/issueshift/internal/web/response"
	"github.com/accmeboot/issueshift/internal/web/validation"
	"net/http"
)

type CreateDTO struct {
	Title       string `form:"title" validate:"required"`
	Description string `form:"description"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var DTO CreateDTO
	err := request.ReadJSON(w, r, &DTO)

	validator := validation.NewValidator()

	if ok := validator.Validate(DTO); !ok {
		response.WriteError(w, http.StatusBadRequest, domain.Envelope{"validation_errors": validator.Errors}, nil)
		return
	}

	err = h.TaskService.Create(DTO.Title, DTO.Description)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "internal service error"}, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
