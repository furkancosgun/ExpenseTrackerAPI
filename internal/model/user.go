package model

import "github.com/furkancosgun/expense-tracker-api/internal/common"

type User struct {
	UserId           string
	Email            string
	FirstName        string
	LastName         string
	Password         string
	AccountConfirmed bool
}

func (user *User) Validate() error {
	if user.Email == "" {
		return common.EMAIL_CANT_BE_EMPTY
	}
	if user.FirstName == "" {
		return common.FIRST_NAME_CANT_BE_EMPTY
	}
	if user.LastName == "" {
		return common.LAST_NAME_CANT_BE_EMPTY
	}
	if user.Password == "" {
		return common.PASSWORD_CANT_BE_EMPTY
	}
	return nil
}
