package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/phayes/freeport"
	pb "ssh-microservice/api"
)

func (c *controller) FreePort(_ context.Context, _ *empty.Empty) (*pb.Port, error) {
	if port, err := freeport.GetFreePort(); err != nil {
		return nil, err
	} else {
		return &pb.Port{Data: uint32(port)}, nil
	}
}
