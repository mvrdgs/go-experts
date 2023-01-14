package service

import (
	"context"
	"fmt"

	"github.com/mvrdgs/go-experts/chi-grpc/internal/pb"
)

func NewGreetService() *GreetService {
	return &GreetService{}
}

type GreetService struct {
	pb.UnimplementedGreetServer
}

func (s *GreetService) SayHello(ctx context.Context, hr *pb.HelloRequest) (*pb.HelloReply, error) {
	msg := fmt.Sprintf("Hello, %s!!", hr.Name)
	return &pb.HelloReply{Message: msg}, nil
}
