package service

import (
	"time"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/helper"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(user dto.UserLoginRequest) (model.User, error)
	Register(user dto.UserRegisterRequest) error
	VerifyAccount(user dto.UserVerifyAccountRequest) error
	ForgotPassword(userForgotPasswordRequest dto.UserForgotPasswordRequest) error
	ResetPassword(user dto.UserResetPasswordRequest) error
}

type AuthService struct {
	authRepository  repository.IUserRepository
	tokenRepository repository.ITokenRepository
}

// ResetPassword implements IAuthService.
func (service *AuthService) ResetPassword(user dto.UserResetPasswordRequest) error {
	var err error

	err = user.Validate()
	if err != nil {
		return err
	}

	findedUser, err := service.authRepository.GetUserByEmail(user.Email)
	if err != nil {
		return common.USER_NOT_FOUND
	}

	//Check Token Sended?
	tokenModel, err := service.tokenRepository.GetTokenByUserId(findedUser.UserId)
	if err != nil || tokenModel.Token != user.Otp || time.Now().After(tokenModel.ExpiresAt) {
		return common.INVALID_OTP_TOKEN
	}

	//Hash Password
	findedUser.Password = helper.HashPassword(user.Password)

	err = service.authRepository.UpdateUser(findedUser)
	return err
}

// ForgotPassword implements IAuthService.
func (service *AuthService) ForgotPassword(userForgotPasswordRequest dto.UserForgotPasswordRequest) error {
	var err error

	err = userForgotPasswordRequest.Validate()

	findedUser, err := service.authRepository.GetUserByEmail(userForgotPasswordRequest.Email)
	if err != nil {
		return common.USER_NOT_FOUND
	}

	service.SendOtpToken(findedUser.Email)

	return err
}

func NewAuthService(authRepository repository.IUserRepository, tokenRepository repository.ITokenRepository) IAuthService {
	return &AuthService{authRepository: authRepository, tokenRepository: tokenRepository}
}

// VerifyAccount implements IAuthService.
func (service *AuthService) VerifyAccount(user dto.UserVerifyAccountRequest) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	findedUser, err := service.authRepository.GetUserByEmail(user.Email)
	if err != nil {
		return common.USER_NOT_FOUND
	}
	//Check Token Sended?
	tokenModel, err := service.tokenRepository.GetTokenByUserId(findedUser.UserId)
	if err != nil {
		return err
	}

	if tokenModel.Token != user.Otp || time.Now().After(tokenModel.ExpiresAt) {
		return common.INVALID_OTP_TOKEN
	}

	findedUser.AccountConfirmed = true
	service.authRepository.UpdateUser(findedUser)

	return nil
}

// Login implements IUserService.
func (service *AuthService) Login(user dto.UserLoginRequest) (model.User, error) {
	var responseUser model.User

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
	user.Password = helper.HashPassword(user.Password)

	//Convert User Model
	userModel := user.ToUser()
	userModel.UserId = currentUser.UserId

	//User Already Have Override Else Save User To Db
	if currentUser.UserId != "" {
		err = service.authRepository.UpdateUser(userModel)
	} else {
		userModel.UserId = uuid.New().String()
		err = service.authRepository.CreateUser(userModel)
	}
	if err != nil {
		return err
	}

	err = service.SendOtpToken(userModel.Email)

	return err
}

func (service *AuthService) SendOtpToken(email string) error {
	var err error

	findedUser, err := service.authRepository.GetUserByEmail(email)
	if err != nil {
		return err
	}

	otpCode := helper.GenerateOTP(6)

	token := model.Token{
		UserId:    findedUser.UserId,
		Token:     otpCode,
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}
	_, err = service.tokenRepository.GetTokenByUserId(findedUser.UserId)
	if err == nil {
		err = service.tokenRepository.UpdateToken(token)
	} else {
		err = service.tokenRepository.CreateToken(token)
	}
	if err == nil {
		err = NewOtpMailContent(findedUser.Email, token.Token, token.ExpiresAt).Send()
	}
	return err
}
