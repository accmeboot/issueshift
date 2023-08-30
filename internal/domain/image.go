package domain

type Image struct {
	ID        int64
	Name      string
	ImageData string
}

type ImageService interface {
	Create() (int64, error)
	Get(id int64) (*Image, error)
}

type ImageRepository interface {
	Create() (int64, error)
	Get(id int64) (*Image, error)
}
