package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"github.com/accmeboot/issueshift/internal/domain"
	"time"
)

func (p *Provider) Authenticate(token string) (*domain.User, error) {
	hash := sha256.Sum256([]byte(token))

	return p.repository.GetUserFromToken(hash[:])
}

func (p *Provider) CreateToken(userId int64) (*string, error) {
	randomBytes := make([]byte, 24)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	tokenPlain := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(tokenPlain))

	return &tokenPlain, p.repository.CreateToken(userId, hash[:], time.Now().Add(3*24*time.Hour))
}
