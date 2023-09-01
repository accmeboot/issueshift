package domain

import "mime/multipart"

type Image struct {
	ID        int64
	Name      string
	ImageData []byte
}

type ImageService interface {
	Create(file *multipart.File, fileName string) (int64, error)
	Get(id int64) (*Image, error)
}

type ImageRepository interface {
	Create(base64File, fileName string) (int64, error)
	Get(id int64) (*string, *string, error)
}
