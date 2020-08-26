package manage

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"ssh-microservice/app/types"
	"ssh-microservice/app/utils"
	"sync"
)

type ClientManager struct {
	options       map[string]*types.SshOption
	tunnels       map[string]*[]types.TunnelOption
	runtime       map[string]*ssh.Client
	localListener *utils.SyncMapListener
	localConn     *utils.SyncMapConn
	remoteConn    *utils.SyncMapConn
	bufPool       *sync.Pool
}

func NewClientManager(poolSize uint32) *ClientManager {
	c := new(ClientManager)
	c.options = make(map[string]*types.SshOption)
	c.tunnels = make(map[string]*[]types.TunnelOption)
	c.runtime = make(map[string]*ssh.Client)
	c.localListener = utils.NewSyncMapListener()
	c.localConn = utils.NewSyncMapConn()
	c.remoteConn = utils.NewSyncMapConn()
	c.bufPool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, poolSize*1024)
		},
	}
	return c
}

func (c *ClientManager) empty(identity string) error {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		return errors.New("this identity does not exists")
	}
	return nil
}

func (c *ClientManager) GetIdentityCollection() []string {
	var keys []string
	for key := range c.options {
		keys = append(keys, key)
	}
	return keys
}

// Get ssh client information
func (c *ClientManager) GetSshOption(identity string) (option *types.SshOption, err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	option = c.options[identity]
	return
}

func (c *ClientManager) GetRuntime(identity string) (client *ssh.Client, err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	client = c.runtime[identity]
	return
}

func (c *ClientManager) GetTunnelOption(identity string) (option []types.TunnelOption, err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	if c.tunnels[identity] != nil {
		option = *c.tunnels[identity]
	}
	return
}
