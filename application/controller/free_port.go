package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ssh-microservice/api"
)

func (c *controller) FreePort(_ context.Context, _ *empty.Empty) (port *pb.Port, err error) {
	//port, err := freeport.GetFreePort()
	//if err != nil {
	//	return &pb.FreePortResponse{
	//		Error: 1,
	//		Msg:   err.Error(),
	//	}, nil
	//}
	//return &pb.FreePortResponse{
	//	Error: 0,
	//	Data:  uint32(port),
	//}, nil
	return
}
