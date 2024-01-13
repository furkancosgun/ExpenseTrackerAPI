package dto

import (
	"strings"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type UserLoginRequest struct {
	Email    string
	Password string
}

func (user *UserLoginRequest) ToNormalized() {
	user.Email = strings.ToUpper(user.Email)
}

func (user *UserLoginRequest) Validate() error {
	if user.Email == "" {
		return common.EMAIL_CANT_BE_EMPTY
	}
	if user.Password == "" {
		return common.PASSWORD_CANT_BE_EMPTY
	}
	return nil
}

type UserVerifyAccountRequest struct {
	Email string
	Otp   string
}

func (user *UserVerifyAccountRequest) ToNormalized() {
	user.Email = strings.ToUpper(user.Email)
}
func (user *UserVerifyAccountRequest) Validate() error {
	if user.Email == "" {
		return common.EMAIL_CANT_BE_EMPTY
	}
	if user.Otp == "" {
		return common.OTP_CODE_CANT_BE_EMPTY
	}
	return nil
}

type UserRegisterRequest struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

func (user *UserRegisterRequest) ToNormalized() {
	user.Email = strings.ToUpper(user.Email)
	user.FirstName = cases.Title(language.English, cases.Compact).String(user.FirstName)
	user.LastName = strings.ToUpper(user.LastName)
}
func (user *UserRegisterRequest) ToUser() model.User {
	return model.User{
		Email:            user.Email,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Password:         user.Password,
		AccountConfirmed: false,
	}
}

func (user *UserRegisterRequest) Validate() error {
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

type UserLoginResponse struct {
	Token string
}
