package main

import (
	"context"
	"fmt"
	"github.com/igorok-follow/grpc-server/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:65000", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := api.NewAdderClient(conn)
	request := &api.Request{
		X: 5,
		Y: 10,
	}
	fmt.Print("REQ: ", request, "\n\n")
	response, err := client.Add(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	log.Println("ADD: ", response.Result)

	response, err = client.Divide(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	log.Println("DIVIDE: ", response.Result)

	response, err = client.Multiply(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	log.Println("MULTIPLY: ", response.Result)

	response, err = client.Subtract(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	log.Println("SUBTRACT: ", response.Result)
}