package controller

import (
	"context"
	"ssh-microservice/common"
	pb "ssh-microservice/router"
)

func (c *controller) Tunnels(ctx context.Context, params *pb.TunnelsParameter) (*pb.CommonResponse, error) {
	var tunnels []common.TunnelOption
	for _, value := range params.Tunnels {
		tunnels = append(tunnels, common.TunnelOption{
			SrcIp:   value.SrcIp,
			SrcPort: value.SrcPort,
			DstIp:   value.DstIp,
			DstPort: value.DstPort,
		})
	}
	err := c.client.SetTunnels(params.Identity, tunnels)
	if err != nil {
		return nil, err
	}
	return &pb.CommonResponse{
		Error: 0,
		Msg:   "ok",
	}, nil
}
