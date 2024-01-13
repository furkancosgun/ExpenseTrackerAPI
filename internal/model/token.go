package model

import "time"

type Token struct {
	Email     string
	Token     string
	ExpiresAt time.Time
}
