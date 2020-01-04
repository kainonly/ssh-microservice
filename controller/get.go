package controller

import (
	"context"
	pb "ssh-microservice/router"
)

func (c *controller) Get(ctx context.Context, params *pb.GetParameter) (*pb.GetResponse, error) {
	data, err := c.client.Get(params.Identity)
	if err != nil {
		return nil, err
	}
	var tunnels []*pb.TunnelOption
	for _, value := range data.Tunnels {
		tunnels = append(tunnels, &pb.TunnelOption{
			SrcIp:   value.SrcIp,
			SrcPort: value.SrcPort,
			DstIp:   value.DstIp,
			DstPort: value.DstPort,
		})
	}
	return &pb.GetResponse{
		Error: 0,
		Data: &pb.Information{
			Identity:  data.Identity,
			Host:      data.Host,
			Port:      data.Port,
			Username:  data.Username,
			Connected: data.Connected,
			Tunnels:   tunnels,
		},
	}, nil
}
