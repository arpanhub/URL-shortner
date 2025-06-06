package main

import (
	"log"
	"net/http"
	"os"
	"github.com/arpanhub/URL-shortner/config"
	"github.com/arpanhub/URL-shortner/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("GIN_MODE") != "release" {
        if err := godotenv.Load(); err != nil {
            log.Println("Warning: .env file not found, relying on environment variables")
        }
    }
	config.ConnectDB()
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	log.Println("URL shortner is running")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "URL shortner is running",
		})
		log.Println("URL shortner is served Request for landing page")
	})

	router.POST("/shorten", handlers.GetShortURL)
	router.GET("/:shortURL", handlers.RedirectURL)
	
	
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server_err := router.Run(":" + port)
	if server_err != nil {
		log.Fatal("Error starting server")
	}
}
