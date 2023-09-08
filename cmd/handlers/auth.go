package handlers

import (
	"errors"
	"github.com/accmeboot/issueshift/internal/domain"
	"net/http"
)

type SignInUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type RegisterUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	Name     string `json:"name" validate:"required"`
	AvatarID int64  `json:"avatar_id"`
}

func (p *Provider) SignUp(w http.ResponseWriter, r *http.Request) {
	var DTO RegisterUserDTO

	err := p.helpers.ReadBody(w, r, &DTO)
	if err != nil {
		p.helpers.SendUnprocessableEntity(w, err)
		return
	}

	validator := p.helpers.NewValidator()
	if ok := validator.Validate(DTO); !ok {
		p.helpers.SendBadRequest(w, validator.Errors, nil)
		return
	}

	err = p.service.CreateUser(DTO.Email, DTO.Name, DTO.Password, &DTO.AvatarID)
	if err != nil {
		var alreadyExists domain.ErrAlreadyExists
		switch {
		case errors.As(err, &alreadyExists):
			p.helpers.SendBadRequest(w, domain.Error{"credentials": "This credentials are not available try again"}, err)
			return
		default:
			p.helpers.SendServerError(w, err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (p *Provider) SignIn(w http.ResponseWriter, r *http.Request) {
	var DTO SignInUserDTO

	err := p.helpers.ReadBody(w, r, &DTO)
	if err != nil {
		p.helpers.SendUnprocessableEntity(w, err)
		return
	}

	validator := p.helpers.NewValidator()

	if ok := validator.Validate(DTO); !ok {
		p.helpers.SendBadRequest(w, validator.Errors, nil)
		return
	}

	user, err := p.service.GetUser(DTO.Email, DTO.Password)
	if err != nil {
		var invalidCredentials domain.ErrInvalidCredentials
		switch {
		case errors.As(err, &invalidCredentials):
			p.helpers.SendError(w, http.StatusNotFound, domain.Error{"credentials": "invalid credentials"}, err)
		default:
			p.helpers.SendServerError(w, err)
		}
		return
	}

	token, err := p.service.CreateToken(user.ID)

	if err != nil {
		p.helpers.SendServerError(w, err)
		return
	}

	p.helpers.Send(w, http.StatusOK, domain.Envelope{"token": token})
}
