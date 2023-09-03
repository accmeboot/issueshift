package web

import (
	"github.com/accmeboot/issueshift/internal/domain"
	"github.com/accmeboot/issueshift/internal/templates"
)

type Handler struct {
	PagesCache   *templates.Cache
	UserService  domain.UserService
	TaskService  domain.TaskService
	ImageService domain.ImageService
	TokenService domain.TokenService
}

func NewHandler(
	pagesCache *templates.Cache,
	userService domain.UserService,
	taskService domain.TaskService,
	imageService domain.ImageService,
	tokenService domain.TokenService,
) *Handler {
	return &Handler{
		PagesCache:   pagesCache,
		UserService:  userService,
		TaskService:  taskService,
		ImageService: imageService,
		TokenService: tokenService,
	}
}
