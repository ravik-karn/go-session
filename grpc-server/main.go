package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const DefaultPort = 5101

func getPort() (int, error) {
	portFromEnv := os.Getenv("APP_PORT")

	appPort, err := strconv.Atoi(portFromEnv)
	if err != nil {
		return 0, errors.New("invalid app port in env")
	}
	return appPort, nil
}

func main() {
	logger := log.New(os.Stdout, "http-server:", log.Ldate| log.Ltime)
	grpclog.SetLogger(logger)

	appPort, err := getPort()
	if err != nil {
		logger.Printf("Using default port %d: %s", DefaultPort, err.Error())
		appPort = DefaultPort
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", appPort))
	if err != nil {
		logger.Fatalf("Failed to listen: %s", err.Error())
	}

	server := grpc.NewServer()
	err = server.Serve(listener)
	if err != nil {
		logger.Fatalf("Server crashed unexpectedly: %s", err.Error())
	}
}
