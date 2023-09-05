package respository

import (
	"database/sql"
	"github.com/accmeboot/issueshift/internal/domain"
)

type Provider struct {
	db *sql.DB
}

func NewProvider(db *sql.DB) *Provider {
	return &Provider{db: db}
}

var _ domain.RepositoryProvider = &Provider{}
