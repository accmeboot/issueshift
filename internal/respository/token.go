package respository

import (
	"context"
	"database/sql"
	"github.com/accmeboot/issueshift/internal/domain"
	"time"
)

type TokenRepository struct {
	DB *sql.DB
}

type Token struct {
	Hash   []byte
	UserID int64
	Expiry time.Time
}

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{
		DB: db,
	}
}

var _ domain.TokenRepository = &TokenRepository{}

func (t TokenRepository) GetUserFromToken(token []byte) (*domain.User, error) {
	query := `SELECT users.id, users.email, users.name, users.password_hash, users.created_at, users.avatar
			  FROM users
			  INNER JOIN tokens
			  ON users.id = tokens.user_id
			  WHERE tokens.hash = $1
			  AND tokens.expiry > $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user domain.User
	var avatarUrl sql.NullString

	err := t.DB.QueryRowContext(ctx, query, token, time.Now()).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.PasswordHash,
		&user.CreatedAt,
		&avatarUrl,
	)
	if err != nil {
		return nil, domain.ErrNoRecord(err)
	}

	if avatarUrl.Valid {
		user.Avatar = &avatarUrl.String
	}

	return &user, nil
}

func (t TokenRepository) Create(userId int64, hash []byte, expiry time.Time) error {
	query := `INSERT INTO tokens (user_id, hash, expiry) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := t.DB.ExecContext(ctx, query, userId, hash, expiry)
	if err != nil {
		return domain.ErrServer(err)
	}

	return nil
}
