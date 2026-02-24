package users_test

import (
	"context"
	"errors"
	"testing"

	users "github.com/deeep8250/services/Users"
	"github.com/jackc/pgx/v5"
)

type fakeUserRepo struct {
	// Controls what repo returns
	getID   int64
	getHash string
	getErr  error

	createID  int64
	createErr error

	// Track calls
	getCalls    int
	createCalls int
	lastEmail   string
}

func (f *fakeUserRepo) GetUserByEmil(ctx context.Context, email string) (int64, string, error) {
	f.getCalls++
	f.lastEmail = email
	return f.getID, f.getHash, f.getErr
}

func (f *fakeUserRepo) CreateUser(ctx context.Context, email, hashedPassword string) (int64, error) {
	f.createCalls++
	f.lastEmail = email
	return f.createID, f.createErr
}

func TestRegisterUser(t *testing.T) {
	cases := []struct {
		name      string
		email     string
		password  string
		getID     int64
		getErr    error
		createID  int64
		createErr error
		wantErr   bool
	}{
		{
			name:     "success",
			email:    "amit@gmail.com",
			password: "pass123",
			getID:    0,
			getErr:   pgx.ErrNoRows,
			createID: 99,
			wantErr:  false,
		},
		{
			name:     "invalid email",
			email:    "bad-email",
			password: "pass123",
			wantErr:  true,
		},
		{
			name:     "user already exists",
			email:    "amit@gmail.com",
			password: "pass123",
			getID:    55,
			getErr:   nil,
			wantErr:  true,
		},
		{
			name:     "db error",
			email:    "amit@gmail.com",
			password: "pass123",
			getID:    0,
			getErr:   errors.New("db down"),
			wantErr:  true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fr := &fakeUserRepo{
				getID:     tc.getID,
				getErr:    tc.getErr,
				createID:  tc.createID,
				createErr: tc.createErr,
			}

			svc := users.NewUserService(fr)
			_, err := svc.RegisterUser(context.Background(), tc.email, tc.name)
			if tc.wantErr && err == nil {
				t.Fatalf("want error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error : %v", err)
			}

		})
	}
}
