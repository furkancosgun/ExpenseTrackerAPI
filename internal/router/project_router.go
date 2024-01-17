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

type ProjectRouter struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}

func NewProjectRouter(ctx *context.Context, dbPool *pgxpool.Pool) *ProjectRouter {
	return &ProjectRouter{ctx: ctx, dbPool: dbPool}
}

func (route *ProjectRouter) RegisterProjectRoutes(router *mux.Router) {
	ctx := context.Background()

	//Repos
	repository := repository.NewProjectRepository(ctx, route.dbPool)

	//Services
	service := service.NewProjectService(repository)

	//Controllers
	controler := controller.NewProjectController(service)

	//Endpoint assigment
	router.HandleFunc(common.BASE_URL+"project/create", controler.CreateProject).Methods("POST")
	router.HandleFunc(common.BASE_URL+"project/list", controler.GetProjectReportByUserId).Methods("POST")
}
