package client

import (
	"io"
	"net"
	"ssh-microservice/application/common/actions"
	"ssh-microservice/config/options"
	"sync"
)

func (c *Client) Tunnels(identity string, tunnelOptions []options.TunnelOption) (err error) {
	if c.Options.Empty(identity) {
		return NotExists
	}
	if !c.Options.TunnelsIsEmpty(identity) {
		c.closeTunnel(identity)
	}
	c.Options.SetTunnels(identity, tunnelOptions)
	for _, tunnelOption := range tunnelOptions {
		go c.setTunnel(identity, tunnelOption)
	}
	return c.schema.Update(*c.Options.Get(identity))
}

// Multiple tunnel settings
func (c *Client) setTunnel(identity string, option options.TunnelOption) {
	localAddr := actions.GetAddr(option.DstIp, uint(option.DstPort))
	remoteAddr := actions.GetAddr(option.SrcIp, uint(option.SrcPort))
	localListener, err := net.Listen("tcp", localAddr)
	if err != nil {
		return
	}
	c.localListener.Put(identity, localAddr, &localListener)
	for {
		localConn, err := localListener.Accept()
		if err != nil {
			return
		}
		c.localConn.Put(identity, localAddr, &localConn)
		remoteConn, err := c.Clients.Get(identity).Dial("tcp", remoteAddr)
		if err != nil {
			return
		}
		c.remoteConn.Put(identity, localAddr, &remoteConn)
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
		io.Copy(*local, *remote)
	}()
	go func() {
		defer wg.Done()
		if _, err := io.Copy(*remote, *local); err != nil {
			(*local).Close()
			(*remote).Close()
		}
		(*remote).Close()
	}()
	wg.Wait()
}

// Close all running tunnels to which the identity belongs
func (c *Client) closeTunnel(identity string) {
	for _, conn := range c.remoteConn.Lists(identity) {
		(*conn).Close()
	}
	c.remoteConn.Remove(identity)
	for _, conn := range c.localConn.Lists(identity) {
		(*conn).Close()
	}
	c.localConn.Remove(identity)
	for _, listener := range c.localListener.Lists(identity) {
		_ = (*listener).Close()
	}
	c.localListener.Remove(identity)
}
