package models
import "time"

type URL struct {
	ID int `db:"id"`
	LongURL string `db:"long_url"`
	ShortURL string `db:"short_url"`
	CreateAt time.Time `db:"create_at"`
	ExpiresAt time.Time `db:"expired_at"`
}
