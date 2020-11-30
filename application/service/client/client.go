package client

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"ssh-microservice/application/service/client/utils"
	"ssh-microservice/application/service/schema"
	"ssh-microservice/config/options"
)

type Client struct {
	Options       *utils.OptionMap
	Clients       *utils.ClientMap
	keepalive     *utils.CronMap
	localListener *utils.ListenerMap
	localConn     *utils.ConnMap
	remoteConn    *utils.ConnMap
	schema        *schema.Schema
}

var (
	NotExists = errors.New("this identity does not exists")
)

func New(schema *schema.Schema) (client *Client, err error) {
	client = new(Client)
	client.Options = utils.NewOptionMap()
	client.Clients = utils.NewClientMap()
	client.keepalive = utils.NewCronMap()
	client.localListener = utils.NewListenerMap()
	client.localConn = utils.NewConnMap()
	client.remoteConn = utils.NewConnMap()
	client.schema = schema
	var clientOptions []options.ClientOption
	if clientOptions, err = client.schema.Lists(); err != nil {
		return
	}
	for _, option := range clientOptions {
		if err = client.Put(options.ClientOption{
			Identity:   option.Identity,
			Host:       option.Host,
			Port:       option.Port,
			Username:   option.Username,
			Password:   option.Password,
			PrivateKey: option.PrivateKey,
			Passphrase: option.Passphrase,
		}); err != nil {
			return
		}
		if len(option.Tunnels) != 0 {
			if err = client.Tunnels(option.Identity, option.Tunnels); err != nil {
				return
			}
		}
	}
	return
}

func (c *Client) GetOptionAndClient(identity string) (*options.ClientOption, *ssh.Client, error) {
	if identity == "" {
		return nil, nil, NotExists
	}
	return c.Options.Get(identity), c.Clients.Get(identity), nil
}
