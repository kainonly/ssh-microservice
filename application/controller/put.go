package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "ssh-microservice/api"
)

func (c *controller) Put(_ context.Context, option *pb.IOption) (_ *empty.Empty, err error) {
	//privateKey, err := base64.StdEncoding.DecodeString(param.PrivateKey)
	//if err != nil {
	//	return c.response(err)
	//}
	//err = c.manager.Put(
	//	param.Identity,
	//	types.SshOption{
	//		Host:       param.Host,
	//		Port:       param.Port,
	//		Username:   param.Username,
	//		Password:   param.Password,
	//		Key:        privateKey,
	//		PassPhrase: []byte(param.Passphrase),
	//	},
	//)
	//if err != nil {
	//	return c.response(err)
	//}
	//return c.response(nil)
	return
}
