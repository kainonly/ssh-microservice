package controller

import (
	pb "ssh-microservice/api"
	"ssh-microservice/application/common"
)

type controller struct {
	pb.UnimplementedAPIServer
	*common.Dependency
}

func New(dep *common.Dependency) *controller {
	c := new(controller)
	c.Dependency = dep
	return c
}

func (c *controller) find(identity string) (*pb.Data, error) {
	data, client, err := c.Client.GetOptionAndClient(identity)
	if err != nil {
		return nil, err
	}
	tunnels := make([]*pb.Tunnel, len(data.Tunnels))
	for key, val := range data.Tunnels {
		tunnels[key] = &pb.Tunnel{
			SrcIp:   val.SrcIp,
			SrcPort: val.SrcPort,
			DstIp:   val.DstIp,
			DstPort: val.DstPort,
		}
	}
	return &pb.Data{
		Id:        data.Identity,
		Host:      data.Host,
		Port:      data.Port,
		Username:  data.Username,
		Connected: string(client.ClientVersion()),
		Tunnels:   tunnels,
	}, nil
}
