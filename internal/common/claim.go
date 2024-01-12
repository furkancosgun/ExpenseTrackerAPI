package common

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	FirstName string
	LastName  string
	Email     string
	jwt.RegisteredClaims
}
