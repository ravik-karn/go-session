package server

import (
	"context"
	"fmt"
	"log"

	"grpc-server/proto/grpc-server"
)

type SomeServer struct {
	logger log.Logger
}

func New(logger log.Logger) grpcserver.GRPCServerServer {
	return &SomeServer{
		logger: logger,
	}
}

func (G SomeServer) Index(context.Context, *grpcserver.IndexRequest) (*grpcserver.IndexResponse, error) {
	return &grpcserver.IndexResponse{
		Data: "Hello",
	}, nil
}

func (G SomeServer) Number(ctx context.Context, req *grpcserver.NumberRequest) (*grpcserver.NumberResponse, error) {
	return &grpcserver.NumberResponse{
		Number: req.Number,
		Data:   fmt.Sprintf("%d", req.Number),
	}, nil
}
