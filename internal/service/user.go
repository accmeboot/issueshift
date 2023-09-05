package service

import (
	"github.com/accmeboot/issueshift/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (p *Provider) GetUser(email, password string) (*domain.User, error) {
	user, err := p.repository.GetUser(email)

	if err != nil {
		return nil, domain.ErrInvalidCredentials(err)
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		return nil, domain.ErrInvalidCredentials(err)
	}

	return user, nil
}

func (p *Provider) CreateUser(email, name, password string, avatarId *int64) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return domain.ErrServer(err)
	}

	return p.repository.CreateUser(email, name, avatarId, hash)
}
