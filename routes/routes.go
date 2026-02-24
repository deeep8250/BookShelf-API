package routes

import (
	"github.com/deeep8250/auth"
	"github.com/deeep8250/database"
	"github.com/deeep8250/handlers"
	BookHandlers "github.com/deeep8250/handlers/Books"
	"github.com/deeep8250/repository"
	Bookservices "github.com/deeep8250/services/Books"
	UserServices "github.com/deeep8250/services/Users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// dependency injection
	//users
	userRepo := repository.NewUserRepository(database.DB)
	userService := UserServices.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	//books
	booksRepo := repository.NewBookRepository(database.DB)
	bookService := Bookservices.NewBookService(booksRepo)
	bookHandler := BookHandlers.NewBookHandler(bookService)

	authGroup := router.Group("/api")
	authGroup.Use(auth.AuthMiddleware())

	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)
	authGroup.POST("/book", bookHandler.CreateBookHandler)
	authGroup.PATCH("update_book/:id", bookHandler.UpdateBookHandler)
	authGroup.DELETE("/books/:id", bookHandler.DeleteBookHandler)
	authGroup.GET("/books", bookHandler.GetBooksHandler)
}
