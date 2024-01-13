package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwrod string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(passwrod), bcrypt.DefaultCost)
	return string(hashedPassword)
}
