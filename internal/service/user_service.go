package service

import (
	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(user dto.UserLoginRequest) (model.User, error)
	Register(user dto.UserRegisterRequest) error
}

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{repository: repository}
}

// Login implements IUserService.
func (service *UserService) Login(user dto.UserLoginRequest) (model.User, error) {
	var responseUser model.User

	//Check Is Valid?
	err := user.Validate()
	if err != nil {
		return responseUser, err
	}

	user.ToNormalized()

	//Check Already Has Email?
	responseUser, err = service.repository.GetUserByEmail(user.Email)
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
func (service *UserService) Register(user dto.UserRegisterRequest) error {
	user.ToNormalized()

	//Check Is Valid?
	err := user.Validate()
	if err != nil {
		return err
	}

	//Check Email Already Using And Confirmed?
	currentUser, err := service.repository.GetUserByEmail(user.Email)
	if err == nil && currentUser.AccountConfirmed {
		return common.EMAIL_ALREADY_USING
	}

	//Hash Password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	//Conver User Model And Normalized value
	model := user.ToUser()

	//User Already Have Override
	if currentUser.Email != "" {
		return service.repository.UpdateUser(model)
	}

	return service.repository.CreateUser(model)
}
