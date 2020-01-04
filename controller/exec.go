package controller

import (
	"context"
	pb "ssh-microservice/router"
)

func (c *controller) Exec(ctx context.Context, params *pb.ExecParameter) (*pb.ExecResponse, error) {
	output, err := c.client.Exec(params.Identity, params.Bash)
	if err != nil {
		return nil, err
	}
	return &pb.ExecResponse{
		Error: 0,
		Data:  string(output),
	}, nil
}
