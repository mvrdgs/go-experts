package main

import (
	"net"

	"github.com/mvrdgs/go-experts/chi-grpc/internal/pb"
	"github.com/mvrdgs/go-experts/chi-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	helloService := service.NewGreetService()

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServer(grpcServer, helloService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
