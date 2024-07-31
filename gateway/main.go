package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
	common "github.com/tuananh9201/commons"
	pb "github.com/tuananh9201/commons/api"
	"github.com/tuananh9201/omsv2-gateway/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	httpAddr         = common.EnvString("GATEWAY_HTTP_PORT", ":3000")
	orderServiceAddr = flag.String("orderServiceAddr", ":50052", "The server address in the format of host:port")
	addr             = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	name             = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// var opts []grpc.DialOption
	// opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	mux := http.NewServeMux()
	// ConnectOrderService(mux)
	// ConnectUserService(mux)
	ConnectGreetingService()
	log.Println("gateway: starting HTTP server at", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("gateway: HTTP server ListenAndServe: %v", err)
	}
}

func ConnectUserService(mux *http.ServeMux) {
	log.Println("gateway: connecting to user service at", addr)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(*addr, opts...)
	if err != nil {
		log.Fatalf("gateway: grpc.NewClient: %v", err)
	}
	defer conn.Close()

	log.Println("gateway: connected to USER service at", addr)
	client := pb.NewUserServiceClient(conn)
	userHandler := handlers.NewUserHandler(client)
	userHandler.RegisterRoutes(mux)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetListUser(ctx, &pb.GetListUserRequest{})
	if err != nil {
		log.Fatalf("gateway: GetListUser: %v", err)
	}
	log.Println(r)

}

func ConnectOrderService(mux *http.ServeMux) {
	log.Println("gateway: connecting to order service at", *orderServiceAddr)
	conn, err := grpc.NewClient(*orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// conn, err := grpc.Dial("localhost:2001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("gateway: grpc.NewClient: %v", err)
	}
	defer conn.Close()

	log.Println("gateway: connected to order service at", *orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)
	handler := handlers.NewHandler(c)
	handler.RegisterRoutes(mux)
}

func ConnectGreetingService() {
	log.Println("gateway: connecting to greeting service at", addr)
	// var opts []grpc.DialOption
	// opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// conn, err := grpc.NewClient(*addr, opts...)
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("gateway: grpc.NewClient: %v", err)
	}
	defer conn.Close()

	log.Println("gateway: connected to Greeting service at", *addr)
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
