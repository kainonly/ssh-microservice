package client

import (
	"golang.org/x/crypto/ssh"
	"log"
	"ssh-microservice/common"
)

type Client struct {
	options       map[string]*common.ConnectOption
	tunnels       map[string]*[]common.TunnelOption
	runtime       map[string]*ssh.Client
	localListener *safeMapListener
	localConn     *safeMapConn
	remoteConn    *safeMapConn
}

// Inject ssh client service
func InjectClient() *Client {
	sshClient := new(Client)
	sshClient.options = make(map[string]*common.ConnectOption)
	sshClient.tunnels = make(map[string]*[]common.TunnelOption)
	sshClient.runtime = make(map[string]*ssh.Client)
	sshClient.localListener = newSafeMapListener()
	sshClient.localConn = newSafeMapConn()
	sshClient.remoteConn = newSafeMapConn()
	configs, err := common.GetTemporary()
	if err != nil {
		log.Fatalln(err)
	}
	if configs.Connect != nil {
		sshClient.options = configs.Connect
	}
	for identity, option := range configs.Connect {
		err = sshClient.Put(identity, *option)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if configs.Tunnel != nil {
		sshClient.tunnels = configs.Tunnel
	}
	for identity, options := range configs.Tunnel {
		err = sshClient.SetTunnels(identity, *options)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return sshClient
}

// Get Client Options
func (c *Client) GetClientOptions() map[string]*common.ConnectOption {
	return c.options
}
