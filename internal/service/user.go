package service

import (
	"github.com/accmeboot/issueshift/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domain.UserRepository
}

var _ domain.UserService = &UserService{}

func NewUserService(r domain.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) GetUserByCredentials(email, password string) (*domain.User, error) {
	user, err := s.repo.GetByEmail(email)

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		return nil, domain.ErrInvalidCredentials(err)
	}

	return s.repo.GetByEmail(email)
}

func (s *UserService) CreateUser(email, name, password string, avatarId *int64) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return domain.ErrServer(err)
	}

	return s.repo.CreateUser(email, name, avatarId, hash)
}
