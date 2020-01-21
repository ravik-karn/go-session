package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"http-server/handlers"
	"http-server/middlewares"
)

const DefaultPort = 5100

func getPort() (int, error) {
	portFromEnv := os.Getenv("APP_PORT")

	appPort, err := strconv.Atoi(portFromEnv)
	if err != nil {
		return 0, errors.New("invalid app port in env")
	}
	return appPort, nil
}

func main(){
	logger := log.New(os.Stdout, "http-server:", log.Ldate| log.Ltime)

	appPort, err := getPort()
	if err != nil {
		logger.Printf("Using default port %d: %s", DefaultPort, err.Error())
		appPort = DefaultPort
	}

	homeHandler := handlers.HomeHandler(logger)
	numberHandler := handlers.NumberHandler(logger)
	userHandler := handlers.UserHandler(logger)

	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods(http.MethodGet)
	router.HandleFunc("/num/{pageNumber}", numberHandler).Methods(http.MethodGet) //num/1
	router.HandleFunc("/user", userHandler).Methods(http.MethodGet)

	router.Use(middlewares.LoggingMiddleware(logger))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", appPort),
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server crashed unexpectedly: %s", err.Error())
	}
	logger.Printf("Server listening on %d", appPort)
}
