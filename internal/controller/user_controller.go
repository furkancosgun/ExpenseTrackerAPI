package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/helper"
	"github.com/furkancosgun/expense-tracker-api/internal/service"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	service service.IUserService
}

func NewUserController(service service.IUserService) *UserController {
	return &UserController{service: service}
}

func (controller *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var userLoginRequest dto.UserLoginRequest

	//Json Decode
	err := json.NewDecoder(r.Body).Decode(&userLoginRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	//Login Service
	user, err := controller.service.Login(userLoginRequest)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	//Create Claim For Token
	expiresAt := time.Now().Add(time.Hour * 24 * 30) //1 Month
	claim := &common.Claim{
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Email:            user.Email,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expiresAt)},
	}

	//Signed Token
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claim)
	tokenString, _ := token.SignedString([]byte(common.JWT_KEY))

	helper.JsonWriteToResponse(w, dto.UserLoginResponse{Token: tokenString}, http.StatusOK)
}

func (controller *UserController) Register(w http.ResponseWriter, r *http.Request) {
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
