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

type CategoryRouter struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}

func NewCategoryRouter(ctx *context.Context, dbPool *pgxpool.Pool) *CategoryRouter {
	return &CategoryRouter{ctx: ctx, dbPool: dbPool}
}

func (route *CategoryRouter) RegisterCategoryRoutes(router *mux.Router) {
	ctx := context.Background()

	//Repos
	repository := repository.NewCategoryRepository(ctx, route.dbPool)

	//Services
	service := service.NewCategoryService(repository)

	//Controllers
	controler := controller.NewCategoryController(service)

	//Endpoint assigment
	router.HandleFunc(common.BASE_URL+"category/create", controler.CreateCategory).Methods("POST")
	router.HandleFunc(common.BASE_URL+"category/list", controler.GetCategories).Methods("POST")
}
