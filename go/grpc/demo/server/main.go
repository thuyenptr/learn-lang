package main

import (
	"log"
	"net"
	pb "demo/demo"
	"google.golang.org/grpc"
	"context"
)

type server struct {}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "50051")
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve %v", err)
	}
}