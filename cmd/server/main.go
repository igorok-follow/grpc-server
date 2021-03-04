package main

import (
	"github.com/igorok-follow/grpc-server/pkg/adder"
	"github.com/igorok-follow/grpc-server/pkg/api"

	"google.golang.org/grpc"
)

func main()  {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

}
