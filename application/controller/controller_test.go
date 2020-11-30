package controller

import (
	"context"
	"encoding/base64"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	pb "ssh-microservice/api"
	"ssh-microservice/bootstrap"
	"ssh-microservice/config"
	"ssh-microservice/config/options"
	"testing"
	"time"
)

var debug []*options.DebugOption
var client pb.APIClient

func TestMain(m *testing.M) {
	os.Chdir("../../")
	var bs []byte
	var err error
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
	if bs, err = ioutil.ReadFile("./config/debug/key-2.pem"); err != nil {
		log.Fatalln(err)
	}
	debug[1].PrivateKey = base64.StdEncoding.EncodeToString(bs)
	var cfg *config.Config
	if cfg, err = bootstrap.LoadConfiguration(); err != nil {
		log.Fatalln(err)
	}
	var conn *grpc.ClientConn
	if conn, err = grpc.Dial(cfg.Listen, grpc.WithInsecure()); err != nil {
		log.Fatalln(err)
	}
	client = pb.NewAPIClient(conn)
	os.Exit(m.Run())
}

func TestController_Testing(t *testing.T) {
	log.Println(debug[0].PrivateKey)
	response, err := client.Testing(
		context.Background(),
		&pb.Option{
			Host:       debug[0].Host,
			Port:       debug[0].Port,
			Username:   debug[0].Username,
			Password:   debug[0].Password,
			PrivateKey: debug[0].PrivateKey,
			Passphrase: debug[0].Passphrase,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestController_Put(t *testing.T) {
	response, err := client.Put(
		context.Background(),
		&pb.IOption{
			Id: "debug-1",
			Option: &pb.Option{
				Host:       debug[0].Host,
				Port:       debug[0].Port,
				Username:   debug[0].Username,
				Password:   debug[0].Password,
				PrivateKey: debug[0].PrivateKey,
				Passphrase: debug[0].Passphrase,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	response, err = client.Put(
		context.Background(),
		&pb.IOption{
			Id: "debug-2",
			Option: &pb.Option{
				Host:       debug[1].Host,
				Port:       debug[1].Port,
				Username:   debug[1].Username,
				Password:   debug[1].Password,
				PrivateKey: debug[1].PrivateKey,
				Passphrase: debug[1].Passphrase,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestController_Exec(t *testing.T) {
	response, err := client.Exec(
		context.Background(),
		&pb.Bash{
			Id:   "debug-1",
			Bash: "uptime",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(response.Data))
	response, err = client.Exec(
		context.Background(),
		&pb.Bash{
			Id:   "debug-2",
			Bash: "uptime",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(response.Data))
}

func TestController_Get(t *testing.T) {
	response, err := client.Get(
		context.Background(),
		&pb.ID{
			Id: "debug-1",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestController_All(t *testing.T) {
	response, err := client.All(
		context.Background(),
		&empty.Empty{},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestController_Lists(t *testing.T) {
	response, err := client.Lists(
		context.Background(),
		&pb.IDs{
			Ids: []string{"debug-1", "debug-2"},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.Data)
}

func TestController_Tunnels(t *testing.T) {
	time.Sleep(time.Second)
	response, err := client.Tunnels(
		context.Background(),
		&pb.TunnelsOption{
			Id: "debug-1",
			Tunnels: []*pb.Tunnel{
				{
					SrcIp:   "127.0.0.1",
					SrcPort: 9200,
					DstIp:   "127.0.0.1",
					DstPort: 9200,
				},
				{
					SrcIp:   "127.0.0.1",
					SrcPort: 5601,
					DstIp:   "127.0.0.1",
					DstPort: 5601,
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	response, err = client.Tunnels(
		context.Background(),
		&pb.TunnelsOption{
			Id: "debug-2",
			Tunnels: []*pb.Tunnel{
				{
					SrcIp:   "127.0.0.1",
					SrcPort: 80,
					DstIp:   "127.0.0.1",
					DstPort: 80,
				},
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestController_Delete(t *testing.T) {
	response, err := client.Delete(
		context.Background(),
		&pb.ID{
			Id: "debug-1",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	response, err = client.Delete(
		context.Background(),
		&pb.ID{
			Id: "debug-2",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestController_FreePort(t *testing.T) {
	response, err := client.FreePort(
		context.Background(),
		&empty.Empty{},
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.Data)
}
