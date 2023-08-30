package factory

import (
	"database/sql"
	"github.com/accmeboot/issueshift/internal/respository"
	"github.com/accmeboot/issueshift/internal/service"
)

type TokenFactory struct {
	TokenService *service.TokenService
}

func NewTokenFactory(db *sql.DB) *TokenFactory {
	tokenRepository := respository.NewTokenRepository(db)
	tokenService := service.NewTokenService(tokenRepository)

	return &TokenFactory{
		TokenService: tokenService,
	}
}
