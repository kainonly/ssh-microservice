package controller

import (
	"context"
	pb "ssh-microservice/api"
)

func (c *controller) Get(_ context.Context, option *pb.ID) (*pb.Data, error) {
	return c.find(option.Id)
}
