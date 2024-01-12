package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/furkancosgun/expense-tracker-api/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	listenAddr := fmt.Sprintf("%s:%s", host, port)

	router := mux.NewRouter()

	router.Use(middleware.LoggerMiddleware)
	router.Use(middleware.AuthenticationMiddleware)
	log.Infof("Server Starting At: %s", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
