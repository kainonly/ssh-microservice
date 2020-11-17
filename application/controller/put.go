package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ssh-microservice/api"
	"ssh-microservice/config/options"
)

func (c *controller) Put(_ context.Context, iOption *pb.IOption) (_ *empty.Empty, err error) {
	option := iOption.Option
	if err = c.Client.Put(options.ClientOption{
		Identity:   iOption.Id,
		Host:       option.Host,
		Port:       option.Port,
		Username:   option.Username,
		Password:   option.Password,
		PrivateKey: option.PrivateKey,
		Passphrase: option.Passphrase,
	}); err != nil {
		return
	}
	return &empty.Empty{}, nil
}
