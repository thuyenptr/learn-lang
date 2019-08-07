package main

import (
	"time"
	"context"
	"log"
	"google.golang.org/grpc"
	pb "firstapp/firstapp_proto"
)

const (
	address = "localhost:50051"
)

func main() {
	conn,err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	client := pb.NewAccountSystemClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.Ping(ctx, &pb.Ping{Timestamp: int64(time.Now().Second())})
	if err != nil {
		log.Fatalf("could not ping %v", err)
	}

	log.Println(res.Timestamp)
	log.Println(res.SystemName)
	log.Println(res.SystemVersion)


}