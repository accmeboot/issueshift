package respository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/accmeboot/issueshift/internal/_model"
	"github.com/accmeboot/issueshift/internal/domain"
	"github.com/lib/pq"
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
			return nil, domain.ErrNoRecord(err)
		default:
			return nil, domain.ErrServer(err)
		}
	}

	return &domain.User{User: user}, nil
}

func (ur *UserRepository) CreateUser(email, name string, avatarUrl *string, passwordHash []byte) error {
	queries := _model.New(ur.DB)
	url := sql.NullString{Valid: false}

	if avatarUrl != nil {
		url.String = *avatarUrl
		url.Valid = true
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := queries.CreateUser(ctx, _model.CreateUserParams{
		Email:        email,
		Name:         name,
		PasswordHash: passwordHash,
		AvatarUrl:    url,
	})

	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code {
		case "23505":
			return domain.ErrAlreadyExists(err)
		default:
			return domain.ErrServer(err)
		}
	}

	return domain.ErrServer(err)
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}
