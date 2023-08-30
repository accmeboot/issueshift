package main

import (
	"database/sql"
	"github.com/accmeboot/issueshift/internal/api"
	"github.com/accmeboot/issueshift/internal/factory"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	User *api.UserHandler

	Mux *chi.Mux
}

func NewHandler(db *sql.DB) *Handler {
	// Initialization of user components
	userFactory := factory.NewUserFactory(db)

	mux := chi.NewRouter()

	// Common middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	return &Handler{
		User: userFactory.Handler,
		Mux:  mux,
	}
}

func (h *Handler) MapUser() {
	h.Mux.Post("/v1/users/signin", h.User.SignInUser)
	h.Mux.Post("/v1/users/signup", h.User.RegisterUser)
}
