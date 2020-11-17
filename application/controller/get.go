package controller

import (
	"context"
	pb "ssh-microservice/api"
)

func (c *controller) Get(_ context.Context, id *pb.ID) (data *pb.Data, err error) {
	//information, err := c.find(param.Identity)
	//if err != nil {
	//	return c.getErrorResponse(err)
	//}
	//return c.getSuccessResponse(information)
	return
}

//func (c *controller) getErrorResponse(err error) (*pb.GetResponse, error) {
//	return &pb.GetResponse{
//		Error: 1,
//		Msg:   err.Error(),
//	}, nil
//}
//
//func (c *controller) getSuccessResponse(data *pb.Information) (*pb.GetResponse, error) {
//	return &pb.GetResponse{
//		Error: 0,
//		Msg:   "ok",
//		Data:  data,
//	}, nil
//}
