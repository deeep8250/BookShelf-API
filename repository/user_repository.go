package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, email, passwordHash string) (int64, error) {
	var userID int64

	query := `INSERT INTO users (email , password_hash)
	          VALUES ($1,$2)
			  RETURNING id 
	`
	err := r.db.QueryRow(ctx, query, email, passwordHash).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil

}
