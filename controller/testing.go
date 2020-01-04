package controller

import (
	"context"
	pb "ssh-microservice/router"
)

func (c *controller) Testing(ctx context.Context, params *pb.TestingParameter) (*pb.CommonResponse, error) {
	return &pb.CommonResponse{
		Error: 0,
		Msg:   "",
	}, nil
}
