package middlewares

import "github.com/accmeboot/issueshift/internal/domain"

type Provider struct {
	service domain.ServiceProvider
}

func NewProvider(service domain.ServiceProvider) *Provider {
	return &Provider{service: service}
}
