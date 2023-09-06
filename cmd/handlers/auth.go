package handlers

import (
	"errors"
	"github.com/accmeboot/issueshift/cmd/helpers"
	"github.com/accmeboot/issueshift/internal/domain"
	"log"
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
}

func (p *Provider) SignInView(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("Bearer")
	if err == nil {
		w.Header().Set("HX-Redirect", "/")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	p.templates.Render(helpers.RenderDTO{
		Writer:   w,
		Template: "signin.gohtml",
	})
}

func (p *Provider) SignUpView(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("Bearer")
	if err == nil {
		w.Header().Set("HX-Redirect", "/")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	p.templates.Render(helpers.RenderDTO{
		Writer:   w,
		Template: "signup.gohtml",
	})
}

func (p *Provider) Logout(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "Bearer",
		MaxAge: -1,
	})
	w.Header().Set("HX-Redirect", "/signin")
}

func (p *Provider) SignUpForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4 << 20)
	if err != nil {
		p.templates.Render(helpers.RenderDTO{
			Writer:   w,
			Template: "signup.gohtml",
			Name:     "signup_form",
			Data: domain.Envelope{
				"validation": domain.Envelope{
					"avatar": "image exceeds max size of 4mb",
				},
			},
		})
		return
	}
	DTO := RegisterUserDTO{
		Name:     r.Form.Get("name"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("password"),
	}

	validator := p.helpers.NewValidator()
	if ok := validator.Validate(DTO); !ok {
		p.templates.Render(helpers.RenderDTO{
			Writer:   w,
			Template: "signup.gohtml",
			Name:     "signup_form",
			Data: domain.Envelope{
				"validation": validator.Errors,
				"email":      &DTO.Email,
				"name":       &DTO.Name,
			},
		})
		return
	}

	file, header, err := r.FormFile("avatar")
	var id int64
	if err == nil {
		defer func() {
			err = file.Close()
			if err != nil {
				panic(err)
			}
		}()
		id, err = p.service.CreateImage(&file, header.Filename)
		if err != nil {
			p.templates.ServerError(w, err)
			return
		}
	}

	var avatarId *int64 = nil
	if id != 0 {
		avatarId = &id
	}

	err = p.service.CreateUser(DTO.Email, DTO.Name, DTO.Password, avatarId)
	if err != nil {
		var alreadyExists domain.ErrAlreadyExists
		log.Println(err)
		switch {
		case errors.As(err, &alreadyExists):
			p.templates.Render(helpers.RenderDTO{
				Writer:   w,
				Template: "signup.gohtml",
				Name:     "signup_form",
				Data: domain.Envelope{
					"error": "These credentials are not available try again",
					"email": &DTO.Email,
					"name":  &DTO.Name,
				},
			})
		default:
			p.templates.ServerError(w, err)
		}
		return
	}

	w.Header().Set("HX-Redirect", "/signin")
}

func (p *Provider) SignInForm(w http.ResponseWriter, r *http.Request) {
	var DTO SignInUserDTO

	err := p.helpers.ReadJSON(w, r, &DTO)
	if err != nil {
		p.templates.ServerError(w, err)
		return
	}

	validator := p.helpers.NewValidator()

	if ok := validator.Validate(DTO); !ok {
		p.templates.Render(helpers.RenderDTO{
			Writer:   w,
			Template: "signin.gohtml",
			Name:     "signin_form",
			Data: domain.Envelope{
				"validation": validator.Errors,
			},
		})
		return
	}

	user, err := p.service.GetUser(DTO.Email, DTO.Password)
	if err != nil {
		var invalidCredentials domain.ErrInvalidCredentials
		switch {
		case errors.As(err, &invalidCredentials):
			p.templates.Render(helpers.RenderDTO{
				Writer:   w,
				Template: "signin.gohtml",
				Name:     "signin_form",
				Data:     domain.Envelope{"error": "invalid credentials"},
			})
		default:
			p.templates.ServerError(w, err)
		}
		return
	}

	token, err := p.service.CreateToken(user.ID)

	if err != nil {
		p.templates.ServerError(w, err)
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
