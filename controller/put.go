package controller

import (
	"context"
	"encoding/base64"
	"ssh-microservice/common"
	pb "ssh-microservice/router"
)

func (c *controller) Put(ctx context.Context, params *pb.PutParameter) (*pb.CommonResponse, error) {
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
	return &pb.CommonResponse{
		Error: 0,
		Msg:   "ok",
	}, nil
}
