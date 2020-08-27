package actions

import (
	"reflect"
	"ssh-microservice/app/types"
	"testing"
)

func TestAuth(t *testing.T) {
	option := debug[0]
	auth, err := Auth(types.SshOption{
		Host:       option.Host,
		Port:       option.Port,
		Username:   option.Username,
		Password:   option.Password,
		Key:        []byte(option.PrivateKey),
		PassPhrase: []byte(option.Passphrase),
	})
	if err != nil {
		t.Fatal(err)
	}
	if reflect.TypeOf(auth).String() == "[]ssh.AuthMethod" {
		t.Logf("[]ssh.AuthMethod created successfully")
	}
}
