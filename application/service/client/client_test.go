package client

import (
	"encoding/base64"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"ssh-microservice/application/service/schema"
	"ssh-microservice/config/options"
	"testing"
)

var client *Client
var debug []*options.DebugOption

func TestMain(m *testing.M) {
	os.Chdir("../../../")
	var err error
	var bs []byte
	if bs, err = ioutil.ReadFile("./config/debug/config.yml"); err != nil {
		log.Fatalln(err)
	}
	if err = yaml.Unmarshal(bs, &debug); err != nil {
		log.Fatalln(err)
	}
	if bs, err = ioutil.ReadFile("./config/debug/key-1.pem"); err != nil {
		log.Fatalln(err)
	}
	debug[0].PrivateKey = base64.StdEncoding.EncodeToString(bs)
	if client, err = New(schema.New("./config/autoload/")); err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}

func TestClientManager_Put(t *testing.T) {
	option := debug[0]
	err := client.Put(options.ClientOption{
		Identity:   "debug-1",
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
}

func TestClientManager_Exec(t *testing.T) {
	output, err := client.Exec("debug-1", "date")
	if err != nil {
		t.Fatal(err)
	}
	if string(output) != "" {
		t.Log("output:", string(output))
	}
}

func TestClientManager_Tunnels(t *testing.T) {
	err := client.Tunnels("debug-1", []options.TunnelOption{
		{
			SrcIp:   "127.0.0.1",
			SrcPort: 3306,
			DstIp:   "127.0.0.1",
			DstPort: 3306,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientManager_Delete(t *testing.T) {
	err := client.Delete("debug-1")
	if err != nil {
		t.Fatal(err)
	}
}
