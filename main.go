package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/skncvo/gin-app/config"
	"github.com/skncvo/gin-app/models"
	"github.com/skncvo/gin-app/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
