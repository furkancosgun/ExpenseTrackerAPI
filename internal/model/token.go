package model

import (
	"strings"
	"time"
)

type Token struct {
	Email     string
	Token     string
	ExpiresAt time.Time
}

func (token *Token) Normalized() {
	token.Email = strings.ToLower(token.Email)
}
