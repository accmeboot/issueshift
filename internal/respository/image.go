package respository

import (
	"context"
	"database/sql"
	"github.com/accmeboot/issueshift/internal/domain"
	"time"
)

type ImageRepository struct {
	DB *sql.DB
}

func NewImageRepository(db *sql.DB) *ImageRepository {
	return &ImageRepository{DB: db}
}

var _ domain.ImageRepository = &ImageRepository{}

func (ir *ImageRepository) Create(base64File, name string) (int64, error) {
	query := `INSERT INTO images (name, image_data) VALUES ($1, $2) RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int64
	err := ir.DB.QueryRowContext(ctx, query, name, base64File).Scan(&id)
	if err != nil {
		return 0, domain.ErrServer(err)
	}

	return id, nil
}

func (ir *ImageRepository) Get(id int64) (*string, *string, error) {
	query := `SELECT name, image_data FROM images WHERE id = $1 LIMIT 1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var imageName string
	var imageData string

	err := ir.DB.QueryRowContext(ctx, query, id).Scan(&imageName, &imageData)
	if err != nil {
		return nil, nil, domain.ErrNoRecord(err)
	}

	return &imageName, &imageData, nil
}
