package api

import (
	"errors"
	"github.com/accmeboot/issueshift/internal/api/validation"
	"log"
	"net/http"

	"github.com/accmeboot/issueshift/internal/api/response"
	"github.com/accmeboot/issueshift/internal/domain"
	"github.com/accmeboot/issueshift/internal/service"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService *service.UserService
}

type SignInUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type RegisterUserDTO struct {
	SignInUserDTO
	Name string `json:"name" validate:"required"`
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (uh *UserHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	var DTO SignInUserDTO

	validator := validation.NewValidator()

	err := response.ReadJSON(w, r, &DTO)
	if err != nil {
		response.WriteError(w, http.StatusUnprocessableEntity, domain.Envelope{"error": "failed to parse body"}, err)
		return
	}

	if ok := validator.Validate(DTO); !ok {
		log.Println(validator.Errors)
		response.WriteError(w, http.StatusBadRequest, domain.Envelope{"validation_errors": validator.Errors}, nil)
		return
	}

	user, err := uh.userService.GetUserByEmail(DTO.Email)
	if err != nil {
		var noRecord domain.ErrNoRecord
		switch {
		case errors.As(err, &noRecord):
			response.WriteError(w, http.StatusNotFound, domain.Envelope{"error": "invalid credentials"}, err)
		default:
			response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "internal server error"}, err)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(DTO.Password))
	if err != nil {
		response.WriteError(w, http.StatusNotFound, domain.Envelope{"error": "invalid credentials"}, err)
		return
	}

	// TODO: create token and save it to db
	mockToken := []byte("token string")

	response.WriteJSON(w, http.StatusOK, domain.Envelope{"token": mockToken})
}

func (uh *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var DTO RegisterUserDTO
	validator := validation.NewValidator()

	err := response.ReadJSON(w, r, &DTO)
	if err != nil {
		response.WriteError(w, http.StatusUnprocessableEntity, domain.Envelope{"error": "failed to parse body"}, err)
		return
	}

	if ok := validator.Validate(DTO); !ok {
		response.WriteError(w, http.StatusBadRequest, domain.Envelope{"validation_errors": validator.Errors}, nil)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(DTO.Password), 12)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, domain.Envelope{"error": "internal server error"}, err)
		return
	}
	// TODO: Add avatar_url logic
	err = uh.userService.CreateUser(DTO.Email, DTO.Name, nil, hash)
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
