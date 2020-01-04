package controller

import (
	"context"
	pb "ssh-microservice/router"
)

func (c *controller) Lists(ctx context.Context, params *pb.ListsParameter) (*pb.ListsResponse, error) {
	var lists []*pb.Information
	for _, identity := range params.Identity {
		data, err := c.client.Get(identity)
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
		lists = append(lists, &pb.Information{
			Identity:  data.Identity,
			Host:      data.Host,
			Port:      data.Port,
			Username:  data.Username,
			Connected: data.Connected,
			Tunnels:   tunnels,
		})
	}
	return &pb.ListsResponse{
		Error: 0,
		Data:  nil,
	}, nil
}
