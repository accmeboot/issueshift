package main

import (
	"database/sql"
	"github.com/accmeboot/issueshift/internal/factory"
	"github.com/accmeboot/issueshift/internal/web"
	"github.com/accmeboot/issueshift/view"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

type Router struct {
	Handler *web.Handler
	Mux     *chi.Mux
}

func NewRouter(db *sql.DB) *Router {
	handler := factory.CreateHandler(db)

	mux := chi.NewRouter()

	// Common middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	fileServer := http.FileServer(http.FS(view.Files))

	//Files server from embedded var
	mux.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})
	mux.Get("/css/*", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})

	return &Router{
		Handler: handler,
		Mux:     mux,
	}
}

func (h *Router) MapRoutes() {
	h.Mux.Get("/", h.Handler.HomeView)
	h.Mux.Get("/signin", h.Handler.SignInView)
	h.Mux.Post("/signin", h.Handler.SignInForm)
	h.Mux.Get("/error", h.Handler.ErrorView)
}
