package main

import (
	"strings"
	"google.golang.org/grpc"
	"log"
	"time"
	"net"
	pb "firstapp/firstapp_proto"
	"context"
)

type server struct {}

func (s *server) Ping(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
	log.Println("Receive: ", in.Timestamp)
	time.Sleep(time.Duration(2 * time.Second))
	return &pb.Pong{
		Timestamp: int64(time.Now().Second()),
		SystemName: "account",
		SystemVersion: "0.0",
	}, nil
}

func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	return &pb.CreateAccountResponse {
		Status: &pb.Status{Code: 0, Error: ""},
		Data: &pb.Account{
			Acno: "1",
			Uid: 1,
			Cino: "3",
			Memo: "memo",
			CreatedTime: time.Now().UTC().String(), 
			Status: pb.AccountStatus_ENABLED,
			Balance: 0,
		},
	}, nil
}


func (s *server) QueryAccount(ctx context.Context, in *pb.AccountQueryRequest) (*pb.AccountResponse, error) {
	return &pb.AccountResponse {
		Status: &pb.Status{Code: 0, Error: ""},
		Data: &pb.Account{
			Acno: "1",
			Uid: 1,
			Cino: "3",
			Memo: "memo",
			CreatedTime: time.Now().UTC().String(), 
			Status: pb.AccountStatus_ENABLED,
			Balance: 0,
		},
	}, nil
}

const (
	port = ":50051"
)


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Printf("Server listening in port %v", strings.Split(port, ":")[1])
	pb.RegisterAccountSystemServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} 
}