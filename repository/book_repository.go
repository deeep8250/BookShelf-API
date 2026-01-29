package repository

import (
	"context"
	"errors"

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

func (r *BookRepository) UpdateBook(ctx context.Context, bookID int64, userID int64, title *string, author *string, description *string) error {

	query := `
	UPDATE books
	SET
	  title = COALESCE($1, title),
	  author = COALESCE($2, author),
	  description = COALESCE($3, description),
	  updated_at = NOW()
	WHERE id = $4 AND user_id = $5
	`

	cmdTag, err := r.db.Exec(
		ctx,
		query,
		title,
		author,
		description,
		bookID,
		userID,
	)

	if err != nil {
		return err
	}

	// Ownership + existence check
	if cmdTag.RowsAffected() == 0 {
		return errors.New("book not found or unauthorized")
	}

	return nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, userID, bookID int64) error {

	query := `DELETE FROM books WHERE id=$1 AND user_id=$2`

	cmdTag, err := r.db.Exec(ctx, query, bookID, userID)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("book not found or unauthorized")
	}

	return nil

}
