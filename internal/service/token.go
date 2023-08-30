package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"github.com/accmeboot/issueshift/internal/domain"
	"time"
)

type TokenService struct {
	repo domain.TokenRepository
}

var _ domain.TokenService = &TokenService{}

func NewTokenService(r domain.TokenRepository) *TokenService {
	return &TokenService{
		repo: r,
	}
}

func (s *TokenService) Authenticate(token string) (*domain.User, error) {
	hash := sha256.Sum256([]byte(token))

	return s.repo.GetUserFromToken(hash[:])
}

func (s *TokenService) Create(userId int64) (*string, error) {
	randomBytes := make([]byte, 24)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	tokenPlain := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(tokenPlain))

	return &tokenPlain, s.repo.Create(userId, hash[:], time.Now().Add(3*24*time.Hour))
}
