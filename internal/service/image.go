package service

import (
	"encoding/base64"
	"github.com/accmeboot/issueshift/internal/domain"
	"io"
	"mime/multipart"
)

func (p *Provider) CreateImage(file *multipart.File, fileName string) (int64, error) {
	fileBytes, err := io.ReadAll(*file)
	if err != nil {
		return 0, domain.ErrServer(err)
	}

	encodedString := base64.StdEncoding.EncodeToString(fileBytes)

	return p.repository.CreateImage(encodedString, fileName)
}

func (p *Provider) GetImage(id int64) (*domain.Image, error) {
	imageName, imageData, err := p.repository.GetImage(id)
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
