package main

import (
	"context"
	"log"
	"net"

	common "github.com/tuananh9201/commons"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", ":2000")
)

func main() {

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("grpc.NewServer: %v", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer()
	store := NewStore()
	svc := NewService(store)

	NewGRPCHandler(grpcServer)
	svc.CreateOrder(context.Background())

	log.Println("orders: starting gRPC server at", l.Addr())
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("grpcServer.Serve: %v", err)
	}
}
