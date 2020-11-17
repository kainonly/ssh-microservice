package actions

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"ssh-microservice/config/options"
	"testing"
)

var debugs []*options.Debug

func TestMain(m *testing.M) {
	os.Chdir("../../")
	bs, err := ioutil.ReadFile("./config/debug/config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(bs, &debugs)
	if err != nil {
		log.Fatalln(err)
	}
	bs, err = ioutil.ReadFile("./config/debug/key-1.pem")
	if err != nil {
		log.Fatalln(err)
	}
	debugs[0].PrivateKey = string(bs)
	os.Exit(m.Run())
}
