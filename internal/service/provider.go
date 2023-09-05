package service

import "github.com/accmeboot/issueshift/internal/domain"

type Provider struct {
	repository domain.RepositoryProvider
}

func NewProvider(repository domain.RepositoryProvider) *Provider {
	return &Provider{repository: repository}
}

var _ domain.ServiceProvider = &Provider{}
