package adder

import (
	"context"
	"github.com/igorok-follow/grpc-server/pkg/api"
)

type GRPCServer struct {}

func (server *GRPCServer) Add(ctx context.Context, req *api.Request) (*api.Response, error) {
	return &api.Response{Result: req.GetX() + req.GetY()}, nil
}

func (server *GRPCServer) Subtract(ctx context.Context, req *api.Request) (*api.Response, error) {
	return &api.Response{Result: req.GetX() - req.GetY()}, nil
}

func (server *GRPCServer) Divide(ctx context.Context, req *api.Request) (*api.Response, error) {
	return &api.Response{Result: req.GetX() / req.GetY()}, nil
}

func (server *GRPCServer) Multiply(ctx context.Context, req *api.Request) (*api.Response, error) {
	return &api.Response{Result: req.GetX() * req.GetY()}, nil
}
