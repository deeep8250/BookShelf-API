package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository struct {
	db *pgxpool.Pool
}

func NewBookRepository(Db *pgxpool.Pool) *BookRepository {
	return &BookRepository{
		db: Db,
	}
}

func (r *BookRepository) CreateBook(ctx context.Context, userID int64, title, author, description string) (int64, error) {

	var bookID int64

	query := `INSERT INTO books (user_id, title, author, description)
	VALUES ($1,$2,$3,$4)
	RETURNING id
	`
	err := r.db.QueryRow(ctx, query, userID, title, author, description).Scan(&bookID)
	if err != nil {
		return 0, err
	}
	return bookID, nil

}
