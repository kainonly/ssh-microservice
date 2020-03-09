package controller

import (
	"context"
	"encoding/base64"
	"ssh-microservice/client"
	"ssh-microservice/common"
	pb "ssh-microservice/router"
)

type controller struct {
	pb.UnimplementedRouterServer
	client client.Client
}

func New(client *client.Client) *controller {
	c := new(controller)
	c.client = *client
	return c
}

func (c *controller) Testing(ctx context.Context, params *pb.TestingParameter) (*pb.Response, error) {
	privateKey, err := base64.StdEncoding.DecodeString(params.PrivateKey)
	if err != nil {
		return nil, err
	}
	cli, err := c.client.Testing(common.ConnectOption{
		Host:       params.Host,
		Port:       params.Port,
		Username:   params.Username,
		Password:   params.Password,
		Key:        privateKey,
		PassPhrase: []byte(params.Passphrase),
	})
	if err != nil {
		return &pb.Response{
			Error: 1,
			Msg:   err.Error(),
		}, nil
	}
	defer cli.Close()
	return &pb.Response{
		Error: 0,
		Msg:   "ok",
	}, nil
}

func (c *controller) Put(ctx context.Context, params *pb.PutParameter) (*pb.Response, error) {
	privateKey, err := base64.StdEncoding.DecodeString(params.PrivateKey)
	if err != nil {
		return nil, err
	}
	err = c.client.Put(params.Identity, common.ConnectOption{
		Host:       params.Host,
		Port:       params.Port,
		Username:   params.Username,
		Password:   params.Password,
		Key:        privateKey,
		PassPhrase: []byte(params.Passphrase),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Error: 0,
		Msg:   "ok",
	}, nil
}

func (c *controller) Exec(ctx context.Context, params *pb.ExecParameter) (*pb.ExecResponse, error) {
	output, err := c.client.Exec(params.Identity, params.Bash)
	if err != nil {
		return nil, err
	}
	return &pb.ExecResponse{
		Error: 0,
		Data:  string(output),
	}, nil
}

func (c *controller) Delete(ctx context.Context, params *pb.DeleteParameter) (*pb.Response, error) {
	err := c.client.Delete(params.Identity)
	if err != nil {
		return &pb.Response{
			Error: 1,
			Msg:   err.Error(),
		}, nil
	} else {
		return &pb.Response{
			Error: 0,
			Msg:   "ok",
		}, nil
	}
}

func (c *controller) Get(ctx context.Context, params *pb.GetParameter) (*pb.GetResponse, error) {
	// TODO:待修改
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

func (c *controller) All(ctx context.Context, params *pb.NoParameter) (*pb.AllResponse, error) {
	var keys []string
	for key := range c.client.GetClientOptions() {
		keys = append(keys, key)
	}
	return &pb.AllResponse{
		Error: 0,
		Data:  keys,
	}, nil
}

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

func (c *controller) Tunnels(ctx context.Context, params *pb.TunnelsParameter) (*pb.Response, error) {
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
	return &pb.Response{
		Error: 0,
		Msg:   "ok",
	}, nil
}
