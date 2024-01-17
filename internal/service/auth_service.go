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
	Login(model model.User) (model.User, error)
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
	return service.authRepository.UpdateUser(findedUser)

}

// Login implements IUserService.
func (service *AuthService) Login(model model.User) (model.User, error) {

	//Check Already Has Email?
	findedUser, err := service.authRepository.GetUserByEmail(model.Email)
	if err != nil {
		return findedUser, common.USER_NOT_FOUND
	}

	//Check Account Confirmed
	if !findedUser.AccountConfirmed {
		return findedUser, common.UN_CONFIRMED_ACCOUNT
	}

	//Check Password Is Same?
	err = bcrypt.CompareHashAndPassword([]byte(findedUser.Password), []byte(model.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return findedUser, common.USER_NOT_FOUND
	}

	return findedUser, nil
}

// Register implements IUserService.
func (service *AuthService) Register(user dto.UserRegisterRequest) error {
	//Check Email Already Using And Confirmed?
	currentUser, err := service.authRepository.GetUserByEmail(user.Email)
	if err == nil && currentUser.AccountConfirmed {
		return common.EMAIL_ALREADY_USING
	}

	//Hash Password
	user.Password = helper.HashPassword(user.Password)

	//Convert User Model
	model := model.User{
		UserId:           currentUser.UserId,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Email:            user.Email,
		Password:         user.Password,
		AccountConfirmed: currentUser.AccountConfirmed,
	}

	//User Already Have Override Else Save User To Db
	if currentUser.UserId != "" {
		err = service.authRepository.UpdateUser(model)
	} else {
		model.UserId = uuid.New().String()
		err = service.authRepository.CreateUser(model)
	}
	if err != nil {
		return err
	}

	err = service.SendOtpToken(model.Email)

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
