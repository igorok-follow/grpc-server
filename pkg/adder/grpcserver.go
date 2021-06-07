package adder

import (
	"context"
	"github.com/igorok-follow/grpc-server/pkg/api"
)

type gRPCServer struct {}

func (server *gRPCServer) Add(ctx context.Context, req *api.AddRequest) (api.AddResponse, error) {
	return api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}