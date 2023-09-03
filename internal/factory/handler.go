package factory

import (
	"database/sql"
	"github.com/accmeboot/issueshift/internal/respository"
	"github.com/accmeboot/issueshift/internal/service"
	"github.com/accmeboot/issueshift/internal/templates"
	"github.com/accmeboot/issueshift/internal/web"
	"log"
)

func CreateHandler(db *sql.DB) *web.Handler {
	templatesCache, err := templates.NewCache()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := respository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	taskRepository := respository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)

	imageRepository := respository.NewImageRepository(db)
	imageService := service.NewImageService(imageRepository)

	tokenRepository := respository.NewTokenRepository(db)
	tokenService := service.NewTokenService(tokenRepository)

	return web.NewHandler(
		templatesCache,
		userService,
		taskService,
		imageService,
		tokenService,
	)
}
