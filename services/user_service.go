package services

import (
	"context"
	"errors"

	"github.com/deeep8250/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, email, password string) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	userID, err := s.userRepo.CreateUser(ctx, email, string(hashedPassword))
	if err != nil {

		return 0, err
	}

	return userID, nil

}
