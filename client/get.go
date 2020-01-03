package client

import (
	"errors"
	"ssh-gRPC/common"
)

type Information struct {
	Identity  string                `json:"identity"`
	Host      string                `json:"host"`
	Port      uint                  `json:"port"`
	Username  string                `json:"username"`
	Connected string                `json:"connected"`
	Tunnels   []common.TunnelOption `json:"tunnels"`
}

// Get ssh client information
func (c *Client) Get(identity string) (content Information, err error) {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		err = errors.New("this identity does not exists")
		return
	}
	option := c.options[identity]
	var tunnels []common.TunnelOption
	if c.tunnels[identity] != nil {
		tunnels = *c.tunnels[identity]
	}
	content = Information{
		Identity:  identity,
		Host:      option.Host,
		Port:      option.Port,
		Username:  option.Username,
		Connected: string(c.runtime[identity].ClientVersion()),
		Tunnels:   tunnels,
	}
	return
}
