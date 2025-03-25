package services
import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
//"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateShortURL generates a random string of length n
func GenerateShortURL() string {
	//this is the source with seeded random value siuu
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	shortUrl := make([]byte,6) //"[][][][][][][][]"
	for i := range shortUrl{
		shortUrl[i] = charset[rng.Intn(len(charset))]
	}	
	return string(shortUrl);
}