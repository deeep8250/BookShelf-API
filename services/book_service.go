package services

import (
	"context"
	"errors"

	"github.com/deeep8250/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(bookRepo *repository.BookRepository) *BookService {
	return &BookService{
		repo: bookRepo,
	}
}

func (s *BookService) CreateBook(ctx context.Context, userID int64, title, author, description string) (int64, error) {
	if title == "" || author == "" || description == "" {
		return 0, errors.New("title is required")
	}

	return s.repo.CreateBook(ctx, userID, title, author, description)
}
