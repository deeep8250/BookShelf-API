package main

import (
	"log"

	"github.com/deeep8250/config"
	"github.com/deeep8250/database"
	"github.com/deeep8250/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()
	database.ConnectDB(cfg)
	router := gin.Default()

	routes.RegisterRoutes(router)

	log.Println("server running on port : ", cfg.Port)
	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal("Failed to connect with the server : ", err)
	}

}
