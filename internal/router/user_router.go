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

type UserRouter struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}

func NewUserRouter(ctx *context.Context, dbPool *pgxpool.Pool) *UserRouter {
	return &UserRouter{ctx: ctx, dbPool: dbPool}
}

func (userRouter UserRouter) RegisterUserRoutes(router *mux.Router) {
	ctx := context.Background()

	respository := repository.NewUserRepository(&ctx, userRouter.dbPool)
	service := service.NewUserService(respository)
	controler := controller.NewUserController(service)

	router.HandleFunc(common.LOGIN_URL, controler.Login).Methods("POST")
	router.HandleFunc(common.REGISTER_URL, controler.Register).Methods("POST")
}
