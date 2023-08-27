package service

import "github.com/accmeboot/issueshift/internal/domain"

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(r domain.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) CreateUser(email, name, avatarUrl string, passwordHash []byte) error {
	return s.repo.CreateUser(email, name, avatarUrl, passwordHash)
}
