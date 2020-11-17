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

//func (c *controller) find(identity string) (information *pb.Information, err error) {
//	sshOption, err := c.manager.GetSshOption(identity)
//	if err != nil {
//		return
//	}
//	client, err := c.manager.GetRuntime(identity)
//	if err != nil {
//		return
//	}
//	tunnelOption, err := c.manager.GetTunnelOption(identity)
//	if err != nil {
//		return
//	}
//	resultTunnelOption := make([]*pb.TunnelOption, len(tunnelOption))
//	for index, option := range tunnelOption {
//		resultTunnelOption[index] = &pb.TunnelOption{
//			SrcIp:   option.SrcIp,
//			SrcPort: option.SrcPort,
//			DstIp:   option.DstIp,
//			DstPort: option.DstPort,
//		}
//	}
//	information = &pb.Information{
//		Identity:  identity,
//		Host:      sshOption.Host,
//		Port:      sshOption.Port,
//		Username:  sshOption.Username,
//		Connected: string(client.ClientVersion()),
//		Tunnels:   resultTunnelOption,
//	}
//	return
//}

//func (c *controller) response(err error) (*pb.Response, error) {
//	if err != nil {
//		return &pb.Response{
//			Error: 1,
//			Msg:   err.Error(),
//		}, nil
//	} else {
//		return &pb.Response{
//			Error: 0,
//			Msg:   "ok",
//		}, nil
//	}
//}
