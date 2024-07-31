package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	// grpcAddr = common.EnvString("ORDER_RGPC_PORT", ":2000")
	port = flag.Int("port", 50052, "The server port")
)

func main() {

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("grpc.NewServer: %v", err)
	}
	// defer l.Close()

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
