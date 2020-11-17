package controller

import (
	"context"
	pb "ssh-microservice/api"
)

func (c *controller) Exec(_ context.Context, option *pb.Bash) (*pb.Output, error) {
	if output, err := c.Client.Exec(option.Id, option.Bash); err != nil {
		return nil, err
	} else {
		return &pb.Output{Data: output}, nil
	}
}
