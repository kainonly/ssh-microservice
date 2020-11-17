package actions

import (
	"reflect"
	pb "ssh-microservice/api"
	"testing"
)

func TestAuth(t *testing.T) {
	option := debugs[0]
	auth, err := Auth(&pb.Option{
		Host:       option.Host,
		Port:       option.Port,
		Username:   option.Username,
		Password:   option.Password,
		PrivateKey: []byte(option.PrivateKey),
		Passphrase: []byte(option.Passphrase),
	})
	if err != nil {
		t.Fatal(err)
	}
	if reflect.TypeOf(auth).String() == "[]ssh.AuthMethod" {
		t.Logf("[]ssh.AuthMethod created successfully")
	}
}
