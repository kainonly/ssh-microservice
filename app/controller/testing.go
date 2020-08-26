package controller

import (
	"context"
	pb "ssh-microservice/router"
)

func (c *controller) Testing(ctx context.Context, req *pb.TestingParameter) (*pb.Response, error) {
	return nil, nil
}
