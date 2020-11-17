package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ssh-microservice/api"
)

func (c *controller) Delete(_ context.Context, option *pb.ID) (_ *empty.Empty, err error) {
	if err = c.Client.Delete(option.Id); err != nil {
		return
	}
	return &empty.Empty{}, nil
}
