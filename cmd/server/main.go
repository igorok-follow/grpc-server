package main

import (
	"github.com/igorok-follow/grpc-server/pkg/adder"
	"github.com/igorok-follow/grpc-server/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":65000")
	if err != nil {
		log.Fatal("Hui tebe v ochko listen pizdanulsa")
	}

	if err = s.Serve(l); err != nil {
		log.Fatal("Hui tebe v ushi serve pizdec prishol")
	}
}