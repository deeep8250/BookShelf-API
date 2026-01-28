package routes

import (
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
	userRepo := repository.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router.POST("/register", userHandler.Register)

}
