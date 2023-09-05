package respository

import (
	"context"
	"github.com/accmeboot/issueshift/internal/domain"
	"time"
)

func (p *Provider) CreateImage(base64File, name string) (int64, error) {
	query := `INSERT INTO images (name, image_data) VALUES ($1, $2) RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int64
	err := p.db.QueryRowContext(ctx, query, name, base64File).Scan(&id)
	if err != nil {
		return 0, domain.ErrServer(err)
	}

	return id, nil
}

func (p *Provider) GetImage(id int64) (*string, *string, error) {
	query := `SELECT name, image_data FROM images WHERE id = $1 LIMIT 1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var imageName string
	var imageData string

	err := p.db.QueryRowContext(ctx, query, id).Scan(&imageName, &imageData)
	if err != nil {
		return nil, nil, domain.ErrNoRecord(err)
	}

	return &imageName, &imageData, nil
}
