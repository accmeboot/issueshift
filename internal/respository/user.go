package respository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/accmeboot/issueshift/internal/domain"
	"github.com/lib/pq"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

type User struct {
	ID           int64
	Email        string
	PasswordHash []byte
	Name         string
	CreatedAt    time.Time
	AvatarID     sql.NullInt64
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

var _ domain.UserRepository = &UserRepository{}

func (ur *UserRepository) GetByEmail(email string) (*domain.User, error) {
	query := `SELECT id, email, name, password_hash, created_at, avatar_id FROM users WHERE email = $1 LIMIT 1`
	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := ur.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.AvatarID,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, domain.ErrNoRecord(err)
		default:
			return nil, domain.ErrServer(err)
		}
	}

	var avatarId *int64 = nil

	if user.AvatarID.Valid {
		avatarId = &user.AvatarID.Int64
	}

	return &domain.User{
		ID:           user.ID,
		Email:        user.Email,
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		AvatarID:     avatarId,
	}, nil
}

func (ur *UserRepository) CreateUser(email, name string, avatarId *int64, passwordHash []byte) error {
	query := `INSERT INTO users (email, name, avatar_id, password_hash) VALUES ($1, $2, $3, $4)`

	url := sql.NullInt64{Valid: false}

	if avatarId != nil {
		url.Int64 = *avatarId
		url.Valid = true
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := ur.DB.ExecContext(ctx, query, email, name, avatarId, passwordHash)

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
