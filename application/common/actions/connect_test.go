package actions

import (
	"ssh-microservice/config/options"
	"testing"
)

func TestConnect(t *testing.T) {
	option := debugs[0]
	client, err := Connect(options.ClientOption{
		Host:       option.Host,
		Port:       option.Port,
		Username:   option.Username,
		Password:   option.Password,
		PrivateKey: option.PrivateKey,
		Passphrase: option.Passphrase,
	})
	if err != nil {
		t.Fatal(err)
	}
	if string(client.ClientVersion()) != "" {
		t.Logf("Successfully connected via go-ssh")
	}
}
