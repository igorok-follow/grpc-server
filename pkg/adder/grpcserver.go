package adder

import (
	"context"
	"github.com/igorok-follow/grpc-server/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {}

// Add ...
func (server *GRPCServer) Add(context *context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse {Res: req.GetX() * req.GetY()}, nil
}
