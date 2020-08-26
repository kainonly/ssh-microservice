package controller

import (
	"context"
	pb "ssh-microservice/router"
)

func (c *controller) Get(ctx context.Context, req *pb.GetParameter) (response *pb.GetResponse, err error) {
	return nil, nil
}
