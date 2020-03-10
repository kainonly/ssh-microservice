package controller

import (
	"context"
	"encoding/base64"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"ssh-microservice/common"
	pb "ssh-microservice/router"
	"testing"
)

var (
	conn  *grpc.ClientConn
	debug *DebugOption
)

type DebugOption struct {
	Host       string `yaml:"host"`
	Port       uint32 `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	PrivateKey string `yaml:"private_key"`
	Passphrase string `yaml:"passphrase"`
}

func TestMain(m *testing.M) {
	in, err := ioutil.ReadFile("../config/config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	cfg := common.AppOption{}
	err = yaml.Unmarshal(in, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	in, err = ioutil.ReadFile("../config/debug.yml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(in, &debug)
	if err != nil {
		log.Fatalln(err)
	}
	in, err = ioutil.ReadFile("../config/key.pem")
	if err != nil {
		log.Fatalln(err)
	}
	debug.PrivateKey = base64.StdEncoding.EncodeToString(in)
	conn, err = grpc.Dial(cfg.Listen, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}

func TestConnect(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Testing(
		context.Background(),
		&pb.TestingParameter{
			Host:       debug.Host,
			Port:       debug.Port,
			Username:   debug.Username,
			Password:   debug.Password,
			PrivateKey: debug.PrivateKey,
			Passphrase: debug.Passphrase,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	}
	println(string(response.Msg))
}

func TestPut(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Put(
		context.Background(),
		&pb.PutParameter{
			Identity:   "test",
			Host:       debug.Host,
			Port:       debug.Port,
			Username:   debug.Username,
			Password:   debug.Password,
			PrivateKey: debug.PrivateKey,
			Passphrase: debug.Passphrase,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	}
	println(string(response.Msg))
}

func TestExec(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Exec(
		context.Background(),
		&pb.ExecParameter{
			Identity: "test",
			Bash:     "uptime",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	}
	println(string(response.Data))
}
