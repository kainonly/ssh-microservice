package controller

import (
	"context"
	pb "ssh-microservice/router"
)

func (c *controller) All(ctx context.Context, req *pb.NoParameter) (*pb.AllResponse, error) {
	return nil, nil
}
