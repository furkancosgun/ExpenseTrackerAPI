package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/furkancosgun/expense-tracker-api/config/database"
	"github.com/furkancosgun/expense-tracker-api/internal/middleware"
	"github.com/furkancosgun/expense-tracker-api/internal/router"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {
	//Load Env
	godotenv.Load("../.env")

	ctx := context.Background()

	//Create Listen Addr
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	listenAddr := fmt.Sprintf("%s:%s", host, port)

	//Create Mux Router
	muxRouter := mux.NewRouter()

	//Add Middlewares
	muxRouter.Use(middleware.ApplicationMiddleware)
	muxRouter.Use(middleware.LoggerMiddleware)
	muxRouter.Use(middleware.AuthenticationMiddleware)

	//Create Db Connection
	dbPool := database.GetPostgresqlConnection(ctx, *database.NewConfig())

	//Assing all routes to mux router
	userRouter := router.NewAuthRouter(&ctx, dbPool)
	userRouter.RegisterUserRoutes(muxRouter)

	//Log it
	log.Infof("Server Starting At: %s", listenAddr)

	http.ListenAndServe(listenAddr, muxRouter)
}
