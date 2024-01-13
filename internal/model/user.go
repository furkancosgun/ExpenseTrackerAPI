package model

import "strings"

type User struct {
	Email            string
	FirstName        string
	LastName         string
	Password         string
	AccountConfirmed bool
}

func (user *User) Normalized() {
	user.Email = strings.ToLower(user.Email)
}
