package manage

import (
	"golang.org/x/crypto/ssh"
	"ssh-microservice/app/types"
	"ssh-microservice/app/utils"
)

type ClientManager struct {
	options       map[string]*types.SshOption
	tunnels       map[string]*[]types.TunnelOption
	runtime       map[string]*ssh.Client
	localListener *utils.SyncMapListener
	localConn     *utils.SyncMapConn
	remoteConn    *utils.SyncMapConn
}

func NewClientManager() *ClientManager {
	c := new(ClientManager)
	c.options = make(map[string]*types.SshOption)
	c.tunnels = make(map[string]*[]types.TunnelOption)
	c.runtime = make(map[string]*ssh.Client)
	c.localListener = utils.NewSyncMapListener()
	c.localConn = utils.NewSyncMapConn()
	c.remoteConn = utils.NewSyncMapConn()
	return c
}
