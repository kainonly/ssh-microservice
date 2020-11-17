package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ssh-microservice/api"
)

func (c *controller) All(_ context.Context, _ *empty.Empty) (ids *pb.IDs, err error) {
	return
}
