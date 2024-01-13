package service

import (
	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/helper"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
)

type ITokenService interface {
	CheckToken(token model.Token) error
	SetToken(email string) error
}

type TokenService struct {
	repository repository.ITokenRepository
}

func NewTokenService(repository repository.ITokenRepository) ITokenService {
	return &TokenService{repository: repository}
}

// GetToken implements ITokenService.
func (service *TokenService) CheckToken(token model.Token) error {
	dbToken, err := service.repository.GetTokenByEmail(token.Email)

	if err != nil || dbToken.Token != token.Token {
		return common.INVALID_OTP_TOKEN
	}

	return nil
}

// SetToken implements ITokenService.
func (service *TokenService) SetToken(email string) error {
	otpCode := helper.GenerateOTP(6)

	token := model.Token{
		Email: email,
		Token: otpCode,
	}

	_, err := service.repository.GetTokenByEmail(email)
	if err == nil {
		return service.repository.UpdateToken(token)
	}
	return service.repository.CreateToken(token)
}
