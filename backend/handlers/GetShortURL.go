package handlers

import (
	"log"
	"net/http"
	// "time"
	// "github.com/arpanhub/URL-shortner/models"
	"github.com/arpanhub/URL-shortner/services"
	"github.com/gin-gonic/gin"
)

func GetShortURL(c *gin.Context) {
	log.Println("IN the URL Handler")
	var request struct{
		LongURL string `json:long_url`
	}
	log.Println("Request is ",request)
	Get_err := c.ShouldBindJSON(&request)
	if Get_err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid request",
		})
		return;
	}else{
		shortUrl := services.GenerateShortURL()
		// creating a url structurre
		// url := models.URL{
		// 	LongURL : request.LongURL,
		// 	ShortURL : shortUrl,
		// 	CreateAt : time.Now(),
		// 	ExpiresAt : time.Now().Add(time.Hour * 24 * 7),
		// }
		c.JSON(http.StatusOK,gin.H{
			"short_url":shortUrl,
		})
	}
}