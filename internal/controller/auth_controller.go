package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/helper"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/service"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	service service.IAuthService
}

func NewAuthController(service service.IAuthService) *AuthController {
	return &AuthController{service: service}
}

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest dto.UserLoginRequest

	//Json Decode
	err := json.NewDecoder(r.Body).Decode(&userLoginRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	model := model.User{
		Email:    userLoginRequest.Email,
		Password: userLoginRequest.Password,
	}

	//Login Service
	user, err := controller.service.Login(model)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	//Create Claim For Token
	expiresAt := time.Now().Add(time.Hour * 24 * 30) //1 Month
	claim := &common.Claim{
		UserId:           user.UserId,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Email:            user.Email,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expiresAt)},
	}

	//Signed Token
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claim)
	tokenString, _ := token.SignedString([]byte(common.JWT_KEY))

	helper.JsonWriteToResponse(w, dto.UserLoginResponse{Token: tokenString,
		FirstName: claim.FirstName,
		LastName:  claim.LastName,
		Email:     claim.Email,
	}, http.StatusOK)
}

func (controller *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var userRegisterRequest dto.UserRegisterRequest

	//Json Decode
	err := json.NewDecoder(r.Body).Decode(&userRegisterRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	//User Register
	err = controller.service.Register(userRegisterRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (controller *AuthController) VerifyAccount(w http.ResponseWriter, r *http.Request) {
	var userVerifyAccounRequest dto.UserVerifyAccountRequest

	//Json Decode
	err := json.NewDecoder(r.Body).Decode(&userVerifyAccounRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = controller.service.VerifyAccount(userVerifyAccounRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (controller *AuthController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var userForgotPasswordRequest dto.UserForgotPasswordRequest

	//Json Decode
	err := json.NewDecoder(r.Body).Decode(&userForgotPasswordRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = controller.service.ForgotPassword(userForgotPasswordRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (controller *AuthController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var userResetPasswordRequest dto.UserResetPasswordRequest

	//Json Decode
	err := json.NewDecoder(r.Body).Decode(&userResetPasswordRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = controller.service.ResetPassword(userResetPasswordRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
