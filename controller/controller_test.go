package controller

import (
	"context"
	"encoding/base64"
	"github.com/sirupsen/logrus"
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
	logrus.Info(response.Msg)
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
	logrus.Info(response.Msg)
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
	logrus.Info(response.Msg)
}

func TestDelete(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Delete(
		context.Background(),
		&pb.DeleteParameter{
			Identity: "test",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	}
	logrus.Info(response.Msg)
}

func TestGet(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Get(
		context.Background(),
		&pb.GetParameter{
			Identity: "test",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	}
	logrus.Info(response.Data)
}

func TestLists(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Lists(
		context.Background(),
		&pb.ListsParameter{
			Identity: []string{"test"},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	}
	logrus.Info(response.Data)
}

func TestAll(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.All(
		context.Background(),
		&pb.NoParameter{},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	}
	logrus.Info(response.Data)
}

func TestTunnels(t *testing.T) {
	defer conn.Close()
	client := pb.NewRouterClient(conn)
	response, err := client.Tunnels(
		context.Background(),
		&pb.TunnelsParameter{
			Identity: "test",
			Tunnels: []*pb.TunnelOption{
				&pb.TunnelOption{
					SrcIp:   "127.0.0.1",
					SrcPort: 80,
					DstIp:   "127.0.0.1",
					DstPort: 8080,
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	}
	logrus.Info(response.Msg)
}
