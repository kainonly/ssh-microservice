package actions

import (
	pb "ssh-microservice/api"
	"testing"
)

func TestConnect(t *testing.T) {
	option := debugs[0]
	client, err := Connect(&pb.Option{
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
	if string(client.ClientVersion()) != "" {
		t.Logf("Successfully connected via go-ssh")
	}
}
