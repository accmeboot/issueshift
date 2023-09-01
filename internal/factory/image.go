package factory

import (
	"database/sql"
	"github.com/accmeboot/issueshift/internal/api"
	"github.com/accmeboot/issueshift/internal/respository"
	"github.com/accmeboot/issueshift/internal/service"
)

type ImageFactory struct {
	Handler *api.ImageHandler
}

func NewImageFactory(db *sql.DB) *ImageFactory {
	imageRepository := respository.NewImageRepository(db)
	imageService := service.NewImageService(imageRepository)

	return &ImageFactory{
		Handler: api.NewImageHandler(imageService),
	}
}
