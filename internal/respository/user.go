package respository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/accmeboot/issueshift/internal/_model"
	"github.com/accmeboot/issueshift/internal/domain"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

var _ domain.UserRepository = &UserRepository{}

func (ur *UserRepository) GetByEmail(email string) (*domain.User, error) {
	queries := _model.New(ur.DB)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err := queries.GetUserByEmail(ctx, email)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, domain.ErrNoRecord
		default:
			return nil, domain.ErrServer
		}
	}

	return &domain.User{User: user}, nil
}

func (ur *UserRepository) CreateUser(email, name, avatarUrl string, passwordHash []byte) error {
	queries := _model.New(ur.DB)
	url := sql.NullString{String: avatarUrl}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := queries.CreateUser(ctx, _model.CreateUserParams{
		Email:        email,
		Name:         name,
		PasswordHash: passwordHash,
		AvatarUrl:    url,
	})

	return err
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}
