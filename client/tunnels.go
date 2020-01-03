package client

import (
	"errors"
	"net"
	"ssh-gRPC/common"
	"sync"
)

// Tunnel setting
func (c *Client) SetTunnels(identity string, options []common.TunnelOption) (err error) {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		err = errors.New("this identity does not exists")
		return
	}
	if c.tunnels[identity] != nil {
		c.closeTunnel(identity)
	}
	c.tunnels[identity] = &options
	for _, tunnel := range options {
		go c.mutilTunnel(identity, tunnel)
	}
	err = common.SetTemporary(common.ConfigOption{
		Connect: c.options,
		Tunnel:  c.tunnels,
	})
	return
}

// Close all running tunnels to which the identity belongs
func (c *Client) closeTunnel(identity string) {
	for _, conn := range c.remoteConn.Map[identity] {
		(*conn).Close()
	}
	c.remoteConn.Clear(identity)
	for _, conn := range c.localConn.Map[identity] {
		(*conn).Close()
	}
	c.localConn.Clear(identity)
	for _, listener := range c.localListener.Map[identity] {
		_ = (*listener).Close()
	}
	c.localListener.Clear(identity)
}

// Multiple tunnel settings
func (c *Client) mutilTunnel(identity string, option common.TunnelOption) {
	localAddr := common.GetAddr(option.DstIp, option.DstPort)
	remoteAddr := common.GetAddr(option.SrcIp, option.SrcPort)
	localListener, err := net.Listen("tcp", localAddr)
	if err != nil {
		println("<" + identity + ">:" + err.Error())
		return
	} else {
		c.localListener.Set(identity, localAddr, &localListener)
	}
	for {
		localConn, err := localListener.Accept()
		if err != nil {
			println("<" + identity + ">:" + err.Error())
			return
		} else {
			c.localConn.Set(identity, localAddr, &localConn)
		}
		remoteConn, err := c.runtime[identity].Dial("tcp", remoteAddr)
		if err != nil {
			println("remote <" + identity + ">:" + err.Error())
			return
		} else {
			c.remoteConn.Set(identity, localAddr, &remoteConn)
		}
		go c.forward(&localConn, &remoteConn)
	}
}

//  Tcp stream forwarding
func (c *Client) forward(local *net.Conn, remote *net.Conn) {
	defer (*local).Close()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		common.Copy(*local, *remote)
	}()
	go func() {
		defer wg.Done()
		if _, err := common.Copy(*remote, *local); err != nil {
			(*local).Close()
			(*remote).Close()
		}
		(*remote).Close()
	}()
	wg.Wait()
}
