package handlers

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/arpanhub/URL-shortner/config"
	"github.com/arpanhub/URL-shortner/models"
	"github.com/arpanhub/URL-shortner/services"
	"github.com/gin-gonic/gin"
)

func GetShortURL(c *gin.Context) {
	log.Println("IN the URL Handler")
	var request struct {
		LongURL   string `json:"long_url"`
		CustomURL string `json:"custom_url"`
	} 

	Get_err := c.ShouldBindJSON(&request)
	if Get_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
		
	}

	log.Println("Request is ", request.LongURL, "Custom Short URL:", request.CustomURL)

	
	var id int
	var long_url string
	var short_url string
	var created_at time.Time
	var expires_at time.Time
	check_err := config.DB.QueryRow(context.Background(),
		`SELECT id, long_url, short_url, created_at, expires_at FROM urls WHERE long_url = $1 AND expires_at > $2`,
		request.LongURL, time.Now()).Scan(&id, &long_url, &short_url, &created_at, &expires_at)
	//Long url is present in the db
	if check_err == nil {
		//no custom string return the present Long url in db
		if request.CustomURL == "" {
			url := models.URL{
				ID:        id,
				LongURL:   long_url,
				ShortURL:  strings.ToLower(short_url),
				CreateAt:  created_at,
				ExpiresAt: expires_at,
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "URL already exists",
				"url":     url,
			})
			return
		} else {
			//user gave the custom string
			user_customURL := request.CustomURL
			var existID int
			//below query checks wether the string in present in db or not
			check_err := config.DB.QueryRow(context.Background(),
				`SELECT id FROM urls WHERE LOWER(short_url) = $1`, user_customURL).Scan(&existID)
			if check_err != nil {
				//this above lines means there exist a short url that why we got some error
				_, update_err := config.DB.Exec(context.Background(),
					`UPDATE urls SET short_url = $1 WHERE id = $2`, strings.ToLower(user_customURL), id)
				if update_err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "Failed to update URL with custom short URL",
					})
					return
				}
				url := models.URL{
					ID:        id,
					LongURL:   long_url,
					ShortURL:  strings.ToLower(user_customURL),
					CreateAt:  created_at,
					ExpiresAt: expires_at,
				}
				c.JSON(http.StatusOK, gin.H{
					"message": "URL updated with custom short URL",
					"url":     url,
				})
				return
			} else {
				c.JSON(http.StatusConflict, gin.H{
					"error": "Custom short URL already exists, please try another",
				})
				return
			}
		}
	} else {
		var url models.URL
		if request.CustomURL == "" {
			random_string := services.GenerateShortURL()
			url = models.URL{
				LongURL:   request.LongURL,
				ShortURL:  strings.ToLower(random_string),
				CreateAt:  time.Now(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
			}
		} else {
			user_customURL := request.CustomURL
			var existID int
			check_err := config.DB.QueryRow(context.Background(),
				`SELECT id FROM urls WHERE LOWER(short_url) = $1`, user_customURL).Scan(&existID)
			if check_err == nil {
				c.JSON(http.StatusConflict, gin.H{
					"error": "Custom short URL already exists, please try another",
				})
				return
			}
			url = models.URL{
				LongURL:   request.LongURL,
				ShortURL:  strings.ToLower(user_customURL),
				CreateAt:  time.Now(),
				ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
			}
		}
		err := config.DB.QueryRow(context.Background(),
			`INSERT INTO urls (long_url, short_url, created_at, expires_at) VALUES($1, $2, $3, $4) RETURNING id`,
			url.LongURL, url.ShortURL, url.CreateAt, url.ExpiresAt).Scan(&url.ID)
		if err != nil {
			log.Println("DB Insert Error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong, please try again",
			})
			return
		}
		c.JSON(http.StatusOK, url)
	}
}
