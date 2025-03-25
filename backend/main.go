package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/arpanhub/URL-shortner/config"
	"github.com/arpanhub/URL-shortner/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func main(){
	port_err := godotenv.Load()
	if port_err != nil{
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	router := gin.Default()
	log.Println("URL shortner is running")
	router.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"URL shortner is running",
		})
		log.Println("URL shortner is served Request for landing page")
	})	

	router.POST("/shorten",handlers.GetShortURL)
	router.GET("/:shortURL",handlers.RedirectURL)
	
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	server_err := router.Run(":" + port)
	if server_err != nil{
		log.Fatal("Error starting server")
	}
}