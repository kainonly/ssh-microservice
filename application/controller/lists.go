package controller

import (
	"context"
	pb "ssh-microservice/api"
)

func (c *controller) Lists(_ context.Context, option *pb.IDs) (*pb.DataLists, error) {
	var err error
	lists := make([]*pb.Data, len(option.Ids))
	for key, val := range option.Ids {
		if lists[key], err = c.find(val); err != nil {
			continue
		}
	}
	return &pb.DataLists{Data: lists}, nil
}
