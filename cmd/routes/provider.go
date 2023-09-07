package routes

import (
	"github.com/accmeboot/issueshift/cmd/handlers"
	"github.com/accmeboot/issueshift/cmd/middlewares"
	"github.com/go-chi/chi/v5"
)

type Provider struct {
	handlers    *handlers.Provider
	middlewares *middlewares.Provider
	Mux         *chi.Mux
}

func NewProvider(handlers *handlers.Provider, middlewares *middlewares.Provider) *Provider {
	provider := &Provider{
		handlers:    handlers,
		middlewares: middlewares,
	}

	provider.CreateRouter()
	provider.MapEndpoints()

	return provider
}
