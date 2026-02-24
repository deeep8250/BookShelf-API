package main

import (
	"log"

	auth "github.com/deeep8250/auth/JWT"
	"github.com/deeep8250/config"
	"github.com/deeep8250/database"
	"github.com/deeep8250/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	database.ConnectDB(cfg)
	auth.InitJWT(cfg.SecretKey)

	router := gin.Default()

	routes.RegisterRoutes(router)

	log.Println("server running on port : ", cfg.Port)
	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal("Failed to connect with the server : ", err)
	}

}
