package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	common "github.com/tuananh9201/commons"
	pb "github.com/tuananh9201/commons/api"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	c := DBConfig{
		Host:     common.EnvString("DB_HOST", "localhost"),
		Port:     common.EnvString("DB_PORT", "5432"),
		User:     common.EnvString("DB_USER", "postgres"),
		Password: common.EnvString("DB_PASSWORD", "postgres"),
		DBName:   common.EnvString("DB_NAME", "item"),
		SSLMode:  common.EnvString("DB_SSL_MODE", "disable"),
	}
	dsn := "host=" + c.Host + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName + " port=" + c.Port + " sslmode=" + c.SSLMode
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println(db)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("grpc.NewServer: %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()
	NewGRPCHandler(s, db)
	pb.RegisterGreeterServer(s, &server{})

	log.Println("users: starting gRPC server at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpcServer.Serve: %v", err)
	}

}
