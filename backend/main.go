package main

import (
	"log"
	"net/http"
	"os"
	"github.com/arpanhub/URL-shortner/config"
	"github.com/arpanhub/URL-shortner/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
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
	
	
	
	// 🚨 Remove fallback port!
	port := os.Getenv("PORT")
	server_err := router.Run(":" + port)
	if server_err != nil {
		log.Fatal("Error starting server")
	}
}
