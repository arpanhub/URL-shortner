package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/arpanhub/URL-shortner/config"
	"github.com/gin-gonic/gin"
)

func RedirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	log.Print("In redirect.go : Short URL is ", shortURL)
	var longUrl string
	var expires_at time.Time
	err := config.DB.QueryRow(context.Background(),
		`SELECT long_url, expires_at FROM urls WHERE LOWER(short_url) = $1`, shortURL).
		Scan(&longUrl, &expires_at)
	log.Print("In redirect.go : Long url found and it is:", longUrl)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "URL not found, make sure you have entered the correct URL",
		})
		return
	}
	if time.Now().After(expires_at) {
		c.JSON(http.StatusGone, gin.H{
			"Message": "URL expired, generate again please",
		})
		return
	}

	_, updateErr := config.DB.Exec(context.Background(),
		`UPDATE urls SET clicks = clicks + 1 WHERE LOWER(short_url) = $1`, shortURL)
	if updateErr != nil {
		log.Println("Unable to increase the click count:", updateErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to increase view count",
		})
		return
	}

	c.Redirect(http.StatusFound, longUrl)
}
