package controller

import (
	"context"
	"encoding/base64"
	"ssh-microservice/common"
	pb "ssh-microservice/router"
)

func (c *controller) Testing(ctx context.Context, params *pb.TestingParameter) (*pb.CommonResponse, error) {
	privateKey, err := base64.StdEncoding.DecodeString(params.PrivateKey)
	if err != nil {
		return nil, err
	}
	sshClient, err := c.client.Testing(common.ConnectOption{
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
	defer sshClient.Close()
	return &pb.CommonResponse{
		Error: 0,
		Msg:   "ok",
	}, nil
}
