package service

import (
	"encoding/base64"
	"github.com/accmeboot/issueshift/internal/domain"
	"io"
	"mime/multipart"
)

type ImageService struct {
	repo domain.ImageRepository
}

var _ domain.ImageService = &ImageService{}

func NewImageService(repo domain.ImageRepository) *ImageService {
	return &ImageService{repo: repo}
}

func (is *ImageService) Create(file *multipart.File, fileName string) (int64, error) {
	fileBytes, err := io.ReadAll(*file)
	if err != nil {
		return 0, domain.ErrServer(err)
	}

	encodedString := base64.StdEncoding.EncodeToString(fileBytes)

	return is.repo.Create(encodedString, fileName)
}

func (is *ImageService) Get(id int64) (*domain.Image, error) {
	imageName, imageData, err := is.repo.Get(id)
	if err != nil {
		return nil, err
	}

	decodedImage, err := base64.StdEncoding.DecodeString(*imageData)
	if err != nil {
		return nil, domain.ErrServer(err)
	}

	return &domain.Image{
		ID:        id,
		Name:      *imageName,
		ImageData: decodedImage,
	}, nil
}