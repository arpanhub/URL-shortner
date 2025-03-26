package models

import "time"

type URL struct {
	ID        int       `db:"id"`
	LongURL   string    `db:"long_url"`
	ShortURL  string    `db:"short_url"`
	CreateAt  time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
	Clicks    int       `db:"clicks"`
}
