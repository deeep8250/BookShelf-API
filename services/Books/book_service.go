package books

import (
	"context"
	"errors"

	"github.com/deeep8250/models"
)

type BookService struct {
	repo BookRepo
}

func NewBookService(bookRepo BookRepo) *BookService {
	return &BookService{
		repo: bookRepo,
	}
}

func (s *BookService) CreateBookHandler(ctx context.Context, userID int64, title, author, description string) (int64, error) {
	if title == "" || author == "" || description == "" {
		return 0, errors.New("title is required")
	}

	return s.repo.CreateBook(ctx, userID, title, author, description)
}

func (s *BookService) GetBoooks(ctx context.Context, userID, limit, offset int) ([]models.Books, error) {

	if limit <= 0 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	if offset < 0 {
		offset = 0
	}

	return s.repo.GetBooksRepo(ctx, userID, limit, offset)

}

func (s *BookService) UpdateBook(ctx context.Context, bookID int64, userID int64, title *string, author *string, description *string) error {

	if title == nil && author == nil && description == nil {
		return errors.New("no fields provided")
	}

	return s.repo.UpdateBook(
		ctx,
		bookID,
		userID,
		title,
		author,
		description,
	)

}

func (s *BookService) DeleteBook(
	ctx context.Context,
	bookID int64,
	userID int64,
) error {

	return s.repo.DeleteBook(ctx, bookID, userID)
}
