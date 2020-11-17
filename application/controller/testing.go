package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/crypto/ssh"
	"log"
	pb "ssh-microservice/api"
	"ssh-microservice/application/common/actions"
	"ssh-microservice/config/options"
)

func (c *controller) Testing(_ context.Context, option *pb.Option) (_ *empty.Empty, err error) {
	log.Println(option)
	var client *ssh.Client
	if client, err = actions.Connect(options.ClientOption{
		Host:       option.Host,
		Port:       option.Port,
		Username:   option.Username,
		Password:   option.Password,
		PrivateKey: option.PrivateKey,
		Passphrase: option.Passphrase,
	}); err != nil {
		return
	}
	defer client.Close()
	var session *ssh.Session
	if session, err = client.NewSession(); err != nil {
		return
	}
	defer session.Close()
	if _, err = session.Output("uptime"); err != nil {
		return
	}
	return &empty.Empty{}, nil
}
