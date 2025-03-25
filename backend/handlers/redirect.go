package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/arpanhub/URL-shortner/config"
	"github.com/gin-gonic/gin"
)
func RedirectURL(c *gin.Context ){
	shortURL := c.Param("shortURL")
	log.Print("In redirect.go : Short URL is ",shortURL)
	var longUrl string
	var expires_at time.Time
	err := config.DB.QueryRow(context.Background(),
		`SELECT long_url,expires_at FROM urls WHERE LOWER(short_url) =$1`,shortURL).
	Scan(&longUrl,&expires_at)
	log.Print("In redirect.go : Long url found and it is:",longUrl)

	if err != nil{
		c.JSON(http.StatusNotFound,gin.H{
			"Message":"URL not found,Make sure you have entered correct URL",
		})
		return
	}
	if time.Now().After(expires_at){
		c.JSON(http.StatusGone,gin.H{
			"Message":"URL expired,Generate again please",
		})
		return
	}
	c.Redirect(http.StatusFound,longUrl)
}