package serv

import (
	"context"

	"github.com/deep-diver/ready/pkg/pbs"
)

type Server struct {
	pbs.UnimplementedDummyServiceServer
}

func (s Server) GetHello(ctx context.Context, in *pbs.HelloRequest) (*pbs.HelloResponse, error) {
	return &pbs.HelloResponse{
		Body: "World",
	}, nil
}
