package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ssh-microservice/api"
)

func (c *controller) Testing(ctx context.Context, option *pb.Option) (_ *empty.Empty, err error) {
	//privateKey, err := base64.StdEncoding.DecodeString(param.PrivateKey)
	//if err != nil {
	//	return c.response(err)
	//}
	//client, err := actions.Connect(types.SshOption{
	//	Host:       param.Host,
	//	Port:       param.Port,
	//	Username:   param.Username,
	//	Password:   param.Password,
	//	Key:        privateKey,
	//	PassPhrase: []byte(param.Passphrase),
	//})
	//if err != nil {
	//	return c.response(err)
	//}
	//defer client.Close()
	//return c.response(nil)
	return
}
