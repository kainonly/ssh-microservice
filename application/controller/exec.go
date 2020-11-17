package controller

import (
	"context"
	pb "ssh-microservice/api"
)

func (c *controller) Exec(_ context.Context, base *pb.Bash) (output *pb.Output, err error) {
	//output, err := c.manager.Exec(param.Identity, param.Bash)
	//if err != nil {
	//	return c.execErrorResponse(err)
	//}
	//return c.execSuccessResponse(output)
	return
}

//func (c *controller) execErrorResponse(err error) (*pb.ExecResponse, error) {
//	return &pb.ExecResponse{
//		Error: 1,
//		Msg:   err.Error(),
//	}, nil
//}
//
//func (c *controller) execSuccessResponse(data []byte) (*pb.ExecResponse, error) {
//	return &pb.ExecResponse{
//		Error: 0,
//		Msg:   "ok",
//		Data:  string(data),
//	}, nil
//}
