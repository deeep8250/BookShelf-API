package users

import "context"

type UserRepo interface {
	GetUserByEmil(ctx context.Context, email string) (int64, string, error)
	CreateUser(ctx context.Context, email, passwordHash string) (int64, error)
}
