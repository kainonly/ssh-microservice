package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ssh-microservice/api"
)

func (c *controller) Delete(_ context.Context, ids *pb.IDs) (_ *empty.Empty, err error) {
	//err := c.manager.Delete(param.Identity)
	//if err != nil {
	//	return c.response(err)
	//}
	return
}
