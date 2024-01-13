package service

import (
	"time"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/helper"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(user dto.UserLoginRequest) (model.User, error)
	Register(user dto.UserRegisterRequest) error
	VerifyAccount(user dto.UserVerifyAccountRequest) error
}

type AuthService struct {
	authRepository  repository.IUserRepository
	tokenRepository repository.ITokenRepository
}

// VerifyAccount implements IAuthService.
func (service *AuthService) VerifyAccount(user dto.UserVerifyAccountRequest) error {
	user.ToNormalized()
	err := user.Validate()
	if err != nil {
		return err
	}

	userModel, err := service.authRepository.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}

	tokenModel, err := service.tokenRepository.GetTokenByEmail(user.Email)
	if err != nil {
		return err
	}

	if tokenModel.Token != user.Otp || time.Now().After(tokenModel.ExpiresAt) {
		return common.INVALID_OTP_TOKEN
	}

	userModel.AccountConfirmed = true
	service.authRepository.UpdateUser(userModel)

	return nil
}

func NewAuthService(authRepository repository.IUserRepository, tokenRepository repository.ITokenRepository) IAuthService {
	return &AuthService{authRepository: authRepository, tokenRepository: tokenRepository}
}

// Login implements IUserService.
func (service *AuthService) Login(user dto.UserLoginRequest) (model.User, error) {
	var responseUser model.User

	user.ToNormalized()

	//Check Is Valid?
	err := user.Validate()
	if err != nil {
		return responseUser, err
	}
	//Check Already Has Email?
	responseUser, err = service.authRepository.GetUserByEmail(user.Email)
	if err != nil {
		return responseUser, common.USER_NOT_FOUND
	}

	//Check Account Confirmed
	if !responseUser.AccountConfirmed {
		return responseUser, common.UN_CONFIRMED_ACCOUNT
	}

	//Check Password Is Same?
	err = bcrypt.CompareHashAndPassword([]byte(responseUser.Password), []byte(user.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return responseUser, common.USER_NOT_FOUND
	}

	return responseUser, nil
}

// Register implements IUserService.
func (service *AuthService) Register(user dto.UserRegisterRequest) error {
	user.ToNormalized()

	//Check Is Valid?
	err := user.Validate()
	if err != nil {
		return err
	}

	//Check Email Already Using And Confirmed?
	currentUser, err := service.authRepository.GetUserByEmail(user.Email)
	if err == nil && currentUser.AccountConfirmed {
		return common.EMAIL_ALREADY_USING
	}

	//Hash Password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	//Conver User Model And Normalized value
	userModel := user.ToUser()

	//User Already Have Override Else Save User To Db
	if currentUser.Email != "" {
		err = service.authRepository.UpdateUser(userModel)
	} else {
		err = service.authRepository.CreateUser(userModel)
	}
	if err != nil {
		return err
	}

	//Generate OTP Code
	otpCode := helper.GenerateOTP(6)

	//Save To Db OTP
	token := model.Token{
		Email:     user.Email,
		Token:     otpCode,
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}
	_, err = service.tokenRepository.GetTokenByEmail(user.Email)
	if err == nil {
		err = service.tokenRepository.UpdateToken(token)
	} else {
		err = service.tokenRepository.CreateToken(token)
	}
	if err == nil {
		err = NewOtpMailContent(token.Email, token.Token, token.ExpiresAt).Send()
	}

	return err
}
