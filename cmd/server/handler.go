package main

import (
	"database/sql"
	"github.com/accmeboot/issueshift/internal/api"
	"github.com/accmeboot/issueshift/internal/factory"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Handler struct {
	User  *api.UserHandler
	Image *api.ImageHandler

	Mux *chi.Mux
}

func NewHandler(db *sql.DB) *Handler {
	// Initialization of user components
	userFactory := factory.NewUserFactory(db)
	imageFactory := factory.NewImageFactory(db)

	mux := chi.NewRouter()

	// Common middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"/", "http://localhost:3000"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return &Handler{
		User:  userFactory.Handler,
		Image: imageFactory.Handler,

		Mux: mux,
	}
}

func (h *Handler) MapUser() {
	h.Mux.Post("/v1/users/signin", h.User.SignInUser)
	h.Mux.Post("/v1/users/signup", h.User.RegisterUser)
}

func (h *Handler) MapImage() {
	h.Mux.Get("/v1/images/{id}", h.Image.GetImage)
	h.Mux.Post("/v1/images", h.Image.CreateImage)
}

func (h *Handler) MapFrontend() {
	h.Mux.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")

		workDir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(workDir, "internal/web/dist"))

		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
	})
}
