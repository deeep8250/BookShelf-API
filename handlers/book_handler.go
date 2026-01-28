package handlers

import (
	"net/http"

	"github.com/deeep8250/services"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(BookService *services.BookService) *BookHandler {
	return &BookHandler{
		bookService: BookService,
	}
}

type CreateBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h *BookHandler) CreateBook(c *gin.Context) {

	var req CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// get the user id from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	bookID, err := h.bookService.CreateBook(c.Request.Context(), userID.(int64), req.Title, req.Author, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"bookID": bookID,
		"msg":    "created!",
	})

}
