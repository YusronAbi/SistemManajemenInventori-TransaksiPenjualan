package grpc

import (
	"google.golang.org/grpc"
)

func ConnectGRPCServer(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
