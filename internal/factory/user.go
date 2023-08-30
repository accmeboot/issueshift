package factory

import (
	"database/sql"
	"github.com/accmeboot/issueshift/internal/api"
	"github.com/accmeboot/issueshift/internal/respository"
	"github.com/accmeboot/issueshift/internal/service"
)

type UserFactory struct {
	Handler *api.UserHandler
}

func NewUserFactory(db *sql.DB) *UserFactory {
	userRepository := respository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	tokenFactory := NewTokenFactory(db)

	return &UserFactory{
		Handler: api.NewUserHandler(userService, tokenFactory.TokenService),
	}
}
