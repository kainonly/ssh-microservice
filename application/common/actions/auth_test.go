package actions

import (
	"reflect"
	"ssh-microservice/config/options"
	"testing"
)

func TestAuth(t *testing.T) {
	option := debugs[0]
	auth, err := Auth(options.ClientOption{
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
	if reflect.TypeOf(auth).String() == "[]ssh.AuthMethod" {
		t.Logf("[]ssh.AuthMethod created successfully")
	}
}
