package dto

import (
	"strings"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
)

type UserLoginRequest struct {
	Email    string
	Password string
}

func (user *UserLoginRequest) Validate() error {
	if user.Email == "" {
		return common.EMAIL_CANT_BE_EMPTY
	}
	if user.Password == "" {
		return common.PASSWORD_CANT_BE_EMPTY
	}
	user.Email = strings.ToLower(user.Email)
	return nil
}

type UserResetPasswordRequest struct {
	Email    string
	Password string
	Otp      string
}

func (user *UserResetPasswordRequest) Validate() error {
	if user.Email == "" {
		return common.EMAIL_CANT_BE_EMPTY
	}
	if user.Password == "" {
		return common.PASSWORD_CANT_BE_EMPTY
	}
	if user.Otp == "" {
		return common.OTP_CODE_CANT_BE_EMPTY
	}
	user.Email = strings.ToLower(user.Email)
	return nil
}

type UserVerifyAccountRequest struct {
	Email string
	Otp   string
}

func (user *UserVerifyAccountRequest) Validate() error {
	if user.Email == "" {
		return common.EMAIL_CANT_BE_EMPTY
	}
	if user.Otp == "" {
		return common.OTP_CODE_CANT_BE_EMPTY
	}
	user.Email = strings.ToLower(user.Email)
	return nil
}

type UserRegisterRequest struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
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
	user.Email = strings.ToLower(user.Email)
	return nil
}

type UserLoginResponse struct {
	FirstName string
	LastName  string
	Email     string
	Token     string
}

type UserForgotPasswordRequest struct {
	Email string
}

func (user *UserForgotPasswordRequest) Validate() error {
	user.Email = strings.TrimSpace(user.Email)
	user.Email = strings.ToLower(user.Email)
	if user.Email == "" {
		return common.EMAIL_CANT_BE_EMPTY
	}
	return nil
}
