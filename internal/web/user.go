package web

import (
	"errors"
	"github.com/accmeboot/issueshift/internal/domain"
	"github.com/accmeboot/issueshift/internal/web/request"
	"github.com/accmeboot/issueshift/internal/web/response"
	"github.com/accmeboot/issueshift/internal/web/validation"
	"net/http"
)

type SignInUserDTO1 struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type RegisterUserDTO struct {
	Name     string `json:"name" validate:"required"`
	AvatarID int64  `json:"avatar_id"`
	SignInUserDTO1
}

//func (h *Handler) SignInUser(w http.ResponseWriter, r *http.Request) {
//	var DTO SignInUserDTO
//
//	err := request.ReadJSON(w, r, &DTO)
//	if err != nil {
//		response.WriteError(w, http.StatusUnprocessableEntity, domain.Envelope{"error": "failed to parse body"}, err)
//		return
//	}
//
//	validator := validation.NewValidator()
//
//	if ok := validator.Validate(DTO); !ok {
//		response.WriteError(w, http.StatusBadRequest, domain.Envelope{"validation_errors": validator.Errors}, nil)
//		return
//	}
//
//	user, err := h.UserService.GetUserByCredentials(DTO.Email, DTO.Password)
//	if err != nil {
//		var invalidCredentials domain.ErrInvalidCredentials
//		switch {
//		case errors.As(err, &invalidCredentials):
//			response.WriteError(w, http.StatusNotFound, domain.Envelope{"error": "invalid credentials"}, err)
//		default:
//			response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "internal server error"}, err)
//		}
//		return
//	}
//
//	token, err := h.TokenService.Create(user.ID)
//
//	if err != nil {
//		response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "internal server error"}, err)
//		return
//	}
//
//	response.WriteJSON(w, http.StatusOK, domain.Envelope{"token": token})
//}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var DTO RegisterUserDTO

	err := request.ReadJSON(w, r, &DTO)
	if err != nil {
		response.WriteError(w, http.StatusUnprocessableEntity, domain.Envelope{"error": "failed to parse body"}, err)
		return
	}

	validator := validation.NewValidator()
	if ok := validator.Validate(DTO); !ok {
		response.WriteError(w, http.StatusBadRequest, domain.Envelope{"validation_errors": validator.Errors}, nil)
		return
	}

	err = h.UserService.CreateUser(DTO.Email, DTO.Name, DTO.Password, &DTO.AvatarID)
	if err != nil {
		var alreadyExists domain.ErrAlreadyExists
		switch {
		case errors.As(err, &alreadyExists):
			response.WriteError(
				w,
				http.StatusBadRequest,
				domain.Envelope{"error": "This credentials are not available try again"},
				err,
			)
			return
		default:
			response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "internal server error"}, err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
