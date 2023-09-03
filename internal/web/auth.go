package web

import (
	"errors"
	"github.com/accmeboot/issueshift/internal/domain"
	"github.com/accmeboot/issueshift/internal/web/request"
	"github.com/accmeboot/issueshift/internal/web/validation"
	"net/http"
)

type SignInUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

func (h *Handler) SignInView(w http.ResponseWriter, _ *http.Request) {
	h.PagesCache.Render(w, http.StatusOK, "signin.gohtml", nil, nil)
}

func (h *Handler) SignInForm(w http.ResponseWriter, r *http.Request) {
	var DTO SignInUserDTO

	err := request.ReadJSON(w, r, &DTO)
	if err != nil {
		h.PagesCache.ServerError(w, err)
		return
	}

	validator := validation.NewValidator()
	fragment := "signin_form"

	if ok := validator.Validate(DTO); !ok {
		// Sending 200 as it is a fragment
		// To enable htmx to swap html with 4.xx codes needs some extensions
		h.PagesCache.Render(w, http.StatusOK, "signin.gohtml", &fragment, validator.Errors)
		return
	}

	user, err := h.UserService.GetUserByCredentials(DTO.Email, DTO.Password)
	if err != nil {
		var invalidCredentials domain.ErrInvalidCredentials
		switch {
		case errors.As(err, &invalidCredentials):
			h.PagesCache.Render(
				w, http.StatusOK,
				"signin.gohtml",
				&fragment,
				domain.Envelope{"notFound": "invalid credentials"},
			)
		default:
			h.PagesCache.ServerError(w, err)
		}
		return
	}

	token, err := h.TokenService.Create(user.ID)

	if err != nil {
		h.PagesCache.ServerError(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer",
		Value:    *token,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	w.Header().Set("HX-Redirect", "/")
}
