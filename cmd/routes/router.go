package routes

import (
	"github.com/accmeboot/issueshift/web"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func (p *Provider) CreateRouter() {
	mux := chi.NewRouter()

	// Common middlewares
	mux.Use(chiMiddleware.Recoverer)
	mux.Use(chiMiddleware.Logger)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	fileServer := http.FileServer(http.FS(web.Files))

	//Files server from embedded var
	mux.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})
	mux.Get("/css/*", func(w http.ResponseWriter, r *http.Request) {
		fileServer.ServeHTTP(w, r)
	})

	p.Mux = mux
}
