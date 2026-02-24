package books

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService BookServices
}

func NewBookHandler(BookService BookServices) *BookHandler {
	return &BookHandler{
		bookService: BookService,
	}
}

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h *BookHandler) CreateBookHandler(c *gin.Context) {

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

	bookID, err := h.bookService.CreateBookHandler(c.Request.Context(), userID.(int64), req.Title, req.Author, req.Description)
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

func (h *BookHandler) GetBooksHandler(c *gin.Context) {

	userIdValue, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	userID := userIdValue.(int64)

	limitStr := c.DefaultQuery("limit", "")
	offsetStr := c.DefaultQuery("offset", "")

	var limit, offset int
	var err error

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid limit",
			})
			return
		}
	}

	if offsetStr != "" {

		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid offset",
			})
			return
		}
	}
	fmt.Println("userID:", userID, "limit:", limit, "offset:", offset)

	books, err := h.bookService.GetBoooks(c.Request.Context(), int(userID), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "hi",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})

}

type UpdateBookRequest struct {
	Title       *string `json:"title"`
	Author      *string `json:"author"`
	Description *string `json:"description"`
}

func (h *BookHandler) UpdateBookHandler(c *gin.Context) {

	bookIdParam := c.Param("id")
	bookID, err := strconv.ParseInt(bookIdParam, 10, 65)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid book id",
		})
		return
	}

	userIdValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	userID := userIdValue.(int64)

	var req UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	err = h.bookService.UpdateBook(
		c.Request.Context(),
		bookID,
		userID,
		req.Title,
		req.Author,
		req.Description,
	)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book updated successfully",
	})

}

func (h *BookHandler) DeleteBookHandler(c *gin.Context) {

	bookIDParam := c.Param("id")

	bookID, err := strconv.ParseInt(bookIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid book id",
		})
		return
	}

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID := userIDValue.(int64)

	err = h.bookService.DeleteBook(
		c.Request.Context(),
		bookID,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "book deleted successfully",
	})
}
