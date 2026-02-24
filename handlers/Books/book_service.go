package books

import (
	"context"

	"github.com/deeep8250/models"
)

type BookServices interface {
	CreateBookHandler(ctx context.Context, userID int64, title, author, description string) (int64, error)
	GetBoooks(ctx context.Context, userID, limit, offset int) ([]models.Books, error)
	UpdateBook(ctx context.Context, bookID int64, userID int64, title *string, author *string, description *string) error
	DeleteBook(ctx context.Context, bookID int64, userID int64) error
}
