package books

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deeep8250/models"
	"github.com/gin-gonic/gin"
)

type fakeService struct {
	BookID      int64
	BookIdError error
	Books       []models.Books
	GetError    error
	UpdateError error
	DeleteError error

	createCall int
	GetCall    int
	UpdateCall int
	DeleteCall int
}

func (fs *fakeService) CreateBookHandler(ctx context.Context, userID int64, title, author, description string) (int64, error) {
	fs.createCall++
	return fs.BookID, fs.BookIdError
}
func (fs *fakeService) GetBoooks(ctx context.Context, userID, limit, offset int) ([]models.Books, error) {
	fs.GetCall++
	return fs.Books, fs.GetError
}
func (fs *fakeService) UpdateBook(ctx context.Context, bookID int64, userID int64, title *string, author *string, description *string) error {
	fs.UpdateCall++
	return fs.UpdateError
}
func (fs *fakeService) DeleteBook(ctx context.Context, bookID int64, userID int64) error {
	fs.DeleteCall++
	return fs.DeleteError
}

func TestCreateBook(t *testing.T) {

	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
	}{
		{
			name:   "success",
			method: http.MethodPost,
			body: `{
				"user_id": 2,
				"title": "Jhingalala",
				"author": "Zuo",
				"description": "Nothing dude"
			}`,
			expectedStatus: http.StatusCreated,
		},
		{
			name:   "missing title",
			method: http.MethodPost,
			body: `{
				"user_id": 2,
				"title": "",
				"author": "Zuo"
			}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "wrong method",
			method:         http.MethodGet,
			body:           "",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "invalid json",
			method:         http.MethodPost,
			body:           `{invalid-json}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			req := httptest.NewRequest(value.method, "/books", bytes.NewBufferString(value.body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			fs := &fakeService{}

			handler := NewBookHandler(fs)

			router := gin.New()
			router.Use(func(c *gin.Context) {
				c.Set("user_id", int64(2))
			})
			router.POST("/books", handler.CreateBookHandler)

			router.ServeHTTP(rec, req)

			if rec.Code != value.expectedStatus {
				t.Fatalf("expected status %d but got %d body=%s", value.expectedStatus, rec.Code, rec.Body.String())
			}
		})

	}

}
func TestGetAllBooks(t *testing.T) {}
func TestUpdateBooks(t *testing.T) {}
func TestDeleteBooks(t *testing.T) {}
