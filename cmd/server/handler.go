package main

import (
	"database/sql"
	"github.com/accmeboot/issueshift/internal/api"
	"github.com/accmeboot/issueshift/internal/api/middleware"
	"github.com/accmeboot/issueshift/internal/factory"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
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
	mux.Use(chiMiddleware.Recoverer)
	mux.Use(chiMiddleware.Logger)
	mux.Use(middleware.ErrorHandler)

	return &Handler{
		User: userFactory.Handler,
		Mux:  mux,
	}
}

func (h *Handler) MapUser() {
	h.Mux.Get("/v1/users", h.User.GetUser)
	h.Mux.Post("/v1/users", h.User.RegisterUser)
}
