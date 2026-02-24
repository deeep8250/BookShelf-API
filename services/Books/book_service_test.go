package books

import (
	"context"
	"errors"
	"testing"

	"github.com/deeep8250/models"
)

type fakeUserRepo struct {
	createId int64
	CreatErr error

	books  []models.Books
	getErr error

	updateErr error
	deleteErr error

	createCalls int
	getCalls    int
	updateCalls int
	deleteCalls int
}

func (f *fakeUserRepo) CreateBook(ctx context.Context, userID int64, title, author, description string) (int64, error) {

	f.createCalls++
	return f.createId, f.CreatErr

}
func (f *fakeUserRepo) GetBooksRepo(ctx context.Context, userID, limit, offset int) ([]models.Books, error) {

	f.getCalls++
	return f.books, f.getErr
}
func (f *fakeUserRepo) UpdateBook(ctx context.Context, bookID int64, userID int64, title *string, author *string, description *string) error {
	f.updateCalls++
	return f.updateErr
}
func (f *fakeUserRepo) DeleteBook(ctx context.Context, userID, bookID int64) error {
	f.deleteCalls++
	return f.deleteErr
}

func TestCreateBookService(t *testing.T) {
	cases := []struct {
		name        string
		title       string
		author      string
		description string
		repoID      int
		repoErr     error
		wantErr     bool
	}{

		{
			name:        "success",
			title:       "Go Mastery",
			author:      "Amit",
			description: "Learn Go",
			repoID:      101,
			repoErr:     nil,
			wantErr:     false,
		},
		{
			name:    "repo error",
			title:   "Go Mastery",
			author:  "Amit",
			repoErr: errors.New("db down"),
			wantErr: true,
		},
	}

	for _, value := range cases {
		t.Run(value.name, func(t *testing.T) {
			fr := &fakeUserRepo{
				createId: int64(value.repoID),
				CreatErr: value.repoErr,
			}
			svc := NewBookService(fr)
			_, err := svc.CreateBookHandler(context.Background(), int64(value.repoID), value.title, value.author, value.description)
			if value.wantErr && err == nil {
				t.Fatalf("expected error but got nil")
			}
			if !value.wantErr && err != nil {
				t.Fatalf("unexepcted error")
			}
		})

	}

}

func TestGetBookService(t *testing.T) {

	cases := []struct {
		id      int
		name    string
		limit   int
		offset  int
		books   []models.Books
		repoErr error
		wantErr bool
	}{

		{
			name:   "success normal",
			id:     1,
			limit:  10,
			offset: 0,
			books: []models.Books{
				{Id: 1, Title: "Go"},
			},
		},
		{
			id:     1,
			name:   "limit default when zero",
			limit:  0,
			offset: 0,
			books:  []models.Books{},
		},
		{
			id:     1,
			name:   "limit capped to 100",
			limit:  200,
			offset: 0,
		},
		{
			id:     1,
			name:   "negative offset corrected",
			limit:  10,
			offset: -5,
		},
		{
			id:      1,
			name:    "repo error",
			limit:   10,
			offset:  0,
			repoErr: errors.New("db down"),
			wantErr: true,
		},
	}

	for _, value := range cases {
		t.Run(value.name, func(t *testing.T) {

			fr := &fakeUserRepo{
				books:  value.books,
				getErr: value.repoErr,
			}

			svc := NewBookService(fr)
			_, err := svc.GetBoooks(context.Background(), value.id, value.limit, value.offset)
			if value.wantErr && err == nil {
				t.Fatalf("want error but got nil")
			}
			if !value.wantErr && err != nil {
				t.Fatalf("unexpected error")
			}

		})
	}

}

func TestDeleteService(t *testing.T) {

	cases := []struct {
		name      string
		userID    int
		bookID    int
		wantError bool
		repoErr   error
	}{
		{
			name:      "success delete",
			userID:    2,
			bookID:    1,
			repoErr:   nil,
			wantError: false,
		},
		{
			name:      "repo error",
			userID:    2,
			bookID:    2,
			repoErr:   errors.New("db failure"),
			wantError: true,
		},
	}

	for _, value := range cases {
		t.Run(value.name, func(t *testing.T) {

			fr := &fakeUserRepo{
				deleteErr: value.repoErr,
			}
			svc := NewBookService(fr)
			err := svc.DeleteBook(context.Background(), int64(value.bookID), int64(value.userID))
			if value.wantError && err == nil {
				t.Fatalf("want error but got nil")
			}
			if !value.wantError && err != nil {
				t.Fatalf("unexpected error")
			}

		})
	}

}
