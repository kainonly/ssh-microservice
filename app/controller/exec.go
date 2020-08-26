package controller

import (
	"context"
	pb "ssh-microservice/router"
)

func (c *controller) Exec(ctx context.Context, param *pb.ExecParameter) (*pb.ExecResponse, error) {
	output, err := c.manager.Exec(param.Identity, param.Bash)
	if err != nil {
		return &pb.ExecResponse{
			Error: 1,
			Msg:   err.Error(),
		}, nil
	}
	return &pb.ExecResponse{
		Error: 0,
		Data:  string(output),
	}, nil
}
