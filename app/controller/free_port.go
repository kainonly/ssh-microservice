package controller

import (
	"context"
	"github.com/phayes/freeport"
	pb "ssh-microservice/router"
)

func (c *controller) FreePort(ctx context.Context, _ *pb.NoParameter) (*pb.FreePortResponse, error) {
	port, err := freeport.GetFreePort()
	if err != nil {
		return &pb.FreePortResponse{
			Error: 1,
			Msg:   err.Error(),
		}, nil
	}
	return &pb.FreePortResponse{
		Error: 0,
		Data:  uint32(port),
	}, nil
}
