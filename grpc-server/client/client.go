package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"grpc-server/proto/grpc-server"
)

func GRPCClient(address string) (int, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithMaxMsgSize(1024*1024*8))
	if err != nil {
		return 0, fmt.Errorf("connection error: %s", err)
	}
	defer conn.Close()
	client := grpcserver.NewGRPCServerClient(conn)
	ctx := context.Background()
	res, err := client.Number(ctx, &grpcserver.NumberRequest{Number: 10})
	if err != nil {
		return 0, fmt.Errorf("server error: %s", err)
	}
	return int(res.Number), nil
}
