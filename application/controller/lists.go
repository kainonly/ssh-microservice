package controller

import (
	"context"
	pb "ssh-microservice/api"
)

func (c *controller) Lists(_ context.Context, ids *pb.IDs) (lists *pb.DataLists, err error) {
	//lists := make([]*pb.Information, len(param.Identity))
	//for index, identity := range param.Identity {
	//	information, err := c.find(identity)
	//	if err != nil {
	//		return c.listsErrorResponse(err)
	//	}
	//	lists[index] = information
	//}
	//return c.listsSuccessResponse(lists)
	return
}

//func (c *controller) listsErrorResponse(err error) (*pb.ListsResponse, error) {
//	return &pb.ListsResponse{
//		Error: 1,
//		Msg:   err.Error(),
//	}, nil
//}
//
//func (c *controller) listsSuccessResponse(data []*pb.Information) (*pb.ListsResponse, error) {
//	return &pb.ListsResponse{
//		Error: 0,
//		Msg:   "ok",
//		Data:  data,
//	}, nil
//}
