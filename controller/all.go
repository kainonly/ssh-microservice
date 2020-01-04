package controller

import (
	"context"
	pb "ssh-microservice/router"
)

func (c *controller) All(ctx context.Context, params *pb.AllParameter) (*pb.AllResponse, error) {
	var keys []string
	for key := range c.client.GetClientOptions() {
		keys = append(keys, key)
	}
	return &pb.AllResponse{
		Error: 0,
		Data:  keys,
	}, nil
}
