package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/nauish/go-auth-docker/models"
	routes "github.com/nauish/go-auth-docker/routes"
)

func main() {
	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	models.Init()

	router.GET("/api/v1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api 1"})
	})

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
