package handlers

import (
	"github.com/accmeboot/issueshift/cmd/helpers"
	"github.com/accmeboot/issueshift/internal/domain"
)

type Provider struct {
	service domain.ServiceProvider
	helpers *helpers.Provider
	pages   *helpers.Cache
}

func NewProvider(service domain.ServiceProvider, pages *helpers.Cache, helpers *helpers.Provider) *Provider {
	return &Provider{service: service, pages: pages, helpers: helpers}
}
