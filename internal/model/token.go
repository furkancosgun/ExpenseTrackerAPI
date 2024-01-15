package model

import (
	"time"
)

type Token struct {
	UserId    string
	Token     string
	ExpiresAt time.Time
}
