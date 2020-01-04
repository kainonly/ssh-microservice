package controller

import (
	"context"
	pb "ssh-microservice/router"
)

func (c *controller) Delete(ctx context.Context, params *pb.DeleteParameter) (*pb.CommonResponse, error) {
	err := c.client.Delete(params.Identity)
	if err != nil {
		return nil, err
	}
	return &pb.CommonResponse{
		Error: 0,
		Msg:   "ok",
	}, nil
}
