package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ssh-microservice/api"
)

func (c *controller) Tunnels(_ context.Context, option *pb.TunnelsOption) (_ *empty.Empty, err error) {
	//tunnelOptions := make([]types.TunnelOption, len(param.Tunnels))
	//for index, option := range param.Tunnels {
	//	tunnelOptions[index] = types.TunnelOption{
	//		SrcIp:   option.SrcIp,
	//		SrcPort: option.SrcPort,
	//		DstIp:   option.DstIp,
	//		DstPort: option.DstPort,
	//	}
	//}
	//err := c.manager.Tunnels(param.Identity, tunnelOptions)
	//if err != nil {
	//	return c.response(err)
	//}
	//return c.response(nil)
	return
}
