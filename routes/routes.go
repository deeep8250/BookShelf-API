package routes

import (
	"github.com/deeep8250/auth"
	"github.com/deeep8250/database"
	"github.com/deeep8250/handlers"
	"github.com/deeep8250/repository"
	"github.com/deeep8250/services"
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
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	//books
	booksRepo := repository.NewBookRepository(database.DB)
	bookService := services.NewBookService(booksRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	authGroup := router.Group("/api")
	authGroup.Use(auth.AuthMiddleware())

	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)
	authGroup.POST("/book", bookHandler.CreateBook)
	authGroup.PATCH("update_book/:id", bookHandler.UpdateBook)
	authGroup.DELETE("/books/:id", bookHandler.DeleteBook)
}
