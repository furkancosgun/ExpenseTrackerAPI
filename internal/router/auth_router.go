package router

import (
	"context"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/controller"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
	"github.com/furkancosgun/expense-tracker-api/internal/service"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthRouter struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}

func NewAuthRouter(ctx *context.Context, dbPool *pgxpool.Pool) *AuthRouter {
	return &AuthRouter{ctx: ctx, dbPool: dbPool}
}

func (authRouter *AuthRouter) RegisterUserRoutes(router *mux.Router) {
	ctx := context.Background()

	//Repos
	userRepository := repository.NewUserRepository(ctx, authRouter.dbPool)
	tokenRepository := repository.NewTokenRepository(ctx, authRouter.dbPool)

	//Services
	service := service.NewAuthService(userRepository, tokenRepository)

	//Controllers
	controler := controller.NewAuthController(service)

	//Endpoint assigment
	router.HandleFunc(common.LOGIN_URL, controler.Login).Methods("POST")
	router.HandleFunc(common.REGISTER_URL, controler.Register).Methods("POST")
	router.HandleFunc(common.VERIFY_ACCOUNT_URL, controler.VerifyAccount).Methods("POST")
}
