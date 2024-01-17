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

type ExpenseRouter struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}

func NewExpenseRouter(ctx *context.Context, dbPool *pgxpool.Pool) *ExpenseRouter {
	return &ExpenseRouter{ctx: ctx, dbPool: dbPool}
}

func (route *ExpenseRouter) RegisterExpensesRoutes(router *mux.Router) {
	ctx := context.Background()

	//Repos
	repository := repository.NewExpenseRepository(ctx, route.dbPool)

	//Services
	service := service.NewExpenseService(repository)

	//Controllers
	controler := controller.NewExpenseController(service)

	//Endpoint assigment
	router.HandleFunc(common.BASE_URL+"expense/create", controler.CreateExpense).Methods("POST")
	router.HandleFunc(common.BASE_URL+"expense/list", controler.GetExpenses).Methods("POST")
}
