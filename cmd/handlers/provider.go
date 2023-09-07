package handlers

import (
	"github.com/accmeboot/issueshift/cmd/helpers"
	"github.com/accmeboot/issueshift/internal/domain"
)

type Provider struct {
	service domain.ServiceProvider
	helpers *helpers.Provider
}

func NewProvider(service domain.ServiceProvider, helpers *helpers.Provider) *Provider {
	return &Provider{service: service, helpers: helpers}
}
