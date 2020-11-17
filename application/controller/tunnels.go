package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ssh-microservice/api"
	"ssh-microservice/config/options"
)

func (c *controller) Tunnels(_ context.Context, option *pb.TunnelsOption) (_ *empty.Empty, err error) {
	tunnels := make([]options.TunnelOption, len(option.Tunnels))
	for key, val := range option.Tunnels {
		tunnels[key] = options.TunnelOption{
			SrcIp:   val.SrcIp,
			SrcPort: val.SrcPort,
			DstIp:   val.DstIp,
			DstPort: val.DstPort,
		}
	}
	if err = c.Client.Tunnels(option.Id, tunnels); err != nil {
		return
	}
	return &empty.Empty{}, nil
}
