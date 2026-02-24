package users

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"

	auth "github.com/deeep8250/auth/JWT"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, email, password string) (int64, error) {

	Email := strings.TrimSpace(email)
	fmt.Println("email is 1 : ", Email)
	_, err := mail.ParseAddress(Email)
	if err != nil {
		fmt.Println("email is : ", Email)
		return 0, errors.New("email address is invalid")
	}

	Password := strings.TrimSpace(password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// veryfying if the user is already exist or not
	id, _, err := s.userRepo.GetUserByEmil(ctx, Email)
	if id != 0 {
		return 0, errors.New("user already exist")
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		return 0, err
	}

	userID, err := s.userRepo.CreateUser(ctx, Email, string(hashedPassword))
	if err != nil {

		return 0, err
	}

	return userID, nil

}

func (s *UserService) LoginUser(ctx context.Context, email, password string) (string, error) {

	userID, storeHash, err := s.userRepo.GetUserByEmil(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(storeHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(userID)
	if err != nil {
		return "", err
	}

	return token, nil

}
