package respository

import (
	"context"
	"database/sql"
	"github.com/accmeboot/issueshift/internal/domain"
	"time"
)

func (p *Provider) GetUserFromToken(token []byte) (*domain.User, error) {
	query := `SELECT users.id, users.email, users.name, users.password_hash, users.created_at, users.avatar_id
			  FROM users
			  INNER JOIN tokens
			  ON users.id = tokens.user_id
			  WHERE tokens.hash = $1
			  AND tokens.expiry > $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user domain.User
	var avatarUrl sql.NullInt64

	err := p.db.QueryRowContext(ctx, query, token, time.Now()).Scan(
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
		user.AvatarID = &avatarUrl.Int64
	}

	return &user, nil
}

func (p *Provider) CreateToken(userId int64, hash []byte, expiry time.Time) error {
	query := `INSERT INTO tokens (user_id, hash, expiry) VALUES ($1, $2, $3)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := p.db.ExecContext(ctx, query, userId, hash, expiry)
	if err != nil {
		return domain.ErrServer(err)
	}

	return nil
}
