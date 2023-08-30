package api

import (
	"errors"
	"github.com/accmeboot/issueshift/internal/api/response"
	"github.com/accmeboot/issueshift/internal/api/validation"
	"github.com/accmeboot/issueshift/internal/domain"
	"log"
	"net/http"
)

type UserHandler struct {
	userService  domain.UserService
	tokenService domain.TokenService
}

type SignInUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type RegisterUserDTO struct {
	Name string `json:"name" validate:"required"`
	SignInUserDTO
}

func NewUserHandler(us domain.UserService, ts domain.TokenService) *UserHandler {
	return &UserHandler{userService: us, tokenService: ts}
}

func (uh *UserHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	var DTO SignInUserDTO

	err := response.ReadJSON(w, r, &DTO)
	if err != nil {
		response.WriteError(w, http.StatusUnprocessableEntity, domain.Envelope{"error": "failed to parse body"}, err)
		return
	}

	validator := validation.NewValidator()

	if ok := validator.Validate(DTO); !ok {
		log.Println(validator.Errors)
		response.WriteError(w, http.StatusBadRequest, domain.Envelope{"validation_errors": validator.Errors}, nil)
		return
	}

	user, err := uh.userService.GetUserByCredentials(DTO.Email, DTO.Password)
	if err != nil {
		var invalidCredentials domain.ErrInvalidCredentials
		switch {
		case errors.As(err, &invalidCredentials):
			response.WriteError(w, http.StatusNotFound, domain.Envelope{"error": "invalid credentials"}, err)
		default:
			response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "internal server error"}, err)
		}
		return
	}

	token, err := uh.tokenService.Create(user.ID)

	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "internal server error"}, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    *token,
		HttpOnly: true,
	})

	response.WriteJSON(w, http.StatusOK, domain.Envelope{"token": token})
}

func (uh *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var DTO RegisterUserDTO

	err := response.ReadJSON(w, r, &DTO)
	if err != nil {
		response.WriteError(w, http.StatusUnprocessableEntity, domain.Envelope{"error": "failed to parse body"}, err)
		return
	}

	validator := validation.NewValidator()
	if ok := validator.Validate(DTO); !ok {
		response.WriteError(w, http.StatusBadRequest, domain.Envelope{"validation_errors": validator.Errors}, nil)
		return
	}

	// TODO: Add avatar_url logic
	err = uh.userService.CreateUser(DTO.Email, DTO.Name, DTO.Password, nil)
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
