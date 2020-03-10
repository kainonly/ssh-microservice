package common

import (
	"encoding/base64"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var debug []*DebugOption

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
	cfg := AppOption{}
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
	in, err = ioutil.ReadFile("../config/key-1.pem")
	if err != nil {
		log.Fatalln(err)
	}
	debug[0].PrivateKey = base64.StdEncoding.EncodeToString(in)
	os.Exit(m.Run())
}

func TestConfig(t *testing.T) {
	if _, err := os.Stat("../config/autoload"); os.IsNotExist(err) {
		os.Mkdir("../config/autoload", os.ModeDir)
	}
}

func TestSaveConfig(t *testing.T) {
	identity := "test"
	option := &ConfigOption{
		Host:       debug[0].Host,
		Port:       debug[0].Port,
		Username:   debug[0].Username,
		Password:   debug[0].Password,
		Key:        debug[0].PrivateKey,
		PassPhrase: debug[0].Passphrase,
		Tunnels: []TunnelOption{
			TunnelOption{
				SrcIp:   "127.0.0.1",
				SrcPort: 9200,
				DstIp:   "127.0.0.1",
				DstPort: 9200,
			},
			TunnelOption{
				SrcIp:   "127.0.0.1",
				SrcPort: 3306,
				DstIp:   "127.0.0.1",
				DstPort: 3306,
			},
		},
	}
	out, err := yaml.Marshal(option)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(
		"../config/autoload/"+identity+".yml",
		out,
		0644,
	)
	if err != nil {
		t.Fatal(err)
	}
}
