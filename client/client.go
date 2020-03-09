package client

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"net"
	"ssh-microservice/common"
	"sync"
)

type Client struct {
	options       map[string]*common.ConnectOption
	tunnels       map[string]*[]common.TunnelOption
	runtime       map[string]*ssh.Client
	localListener *safeMapListener
	localConn     *safeMapConn
	remoteConn    *safeMapConn
}

// Create ssh client service
func Create() *Client {
	client := new(Client)
	client.options = make(map[string]*common.ConnectOption)
	client.tunnels = make(map[string]*[]common.TunnelOption)
	client.runtime = make(map[string]*ssh.Client)
	client.localListener = newSafeMapListener()
	client.localConn = newSafeMapConn()
	client.remoteConn = newSafeMapConn()
	return client
}

// Generate Auth Method
func (c *Client) auth(option common.ConnectOption) (auth []ssh.AuthMethod, err error) {
	if option.Key == nil {
		// Password AuthMethod
		auth = []ssh.AuthMethod{
			ssh.Password(option.Password),
		}
	} else {
		// PrivateKey AuthMethod
		var signer ssh.Signer
		if option.PassPhrase != nil {
			// With Passphrase
			if signer, err = ssh.ParsePrivateKeyWithPassphrase(
				option.Key,
				option.PassPhrase,
			); err != nil {
				return
			}
		} else {
			// Without Passphrase
			if signer, err = ssh.ParsePrivateKey(
				option.Key,
			); err != nil {
				return
			}
		}
		auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
	}
	return
}

// Ssh client connection
func (c *Client) connect(option common.ConnectOption) (client *ssh.Client, err error) {
	auth, err := c.auth(option)
	if err != nil {
		return
	}
	config := ssh.ClientConfig{
		User:            option.Username,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := common.GetAddr(option.Host, uint(option.Port))
	client, err = ssh.Dial("tcp", addr, &config)
	return
}

func (c *Client) empty(identity string) error {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		return errors.New("this identity does not exists")
	}
	return nil
}

// Test ssh client connection
func (c *Client) Testing(option common.ConnectOption) (*ssh.Client, error) {
	return c.connect(option)
}

// Add or modify the ssh client
func (c *Client) Put(identity string, option common.ConnectOption) (err error) {
	if err = c.Delete(identity); err != nil {
		return
	}
	if c.tunnels[identity] != nil {
		c.closeTunnel(identity)
	}
	if c.runtime[identity] != nil {
		c.runtime[identity].Close()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.options[identity] = &option
		c.runtime[identity], err = c.connect(option)
		if err != nil {
			return
		}
	}()
	wg.Wait()
	return
}

// Remotely execute commands via SSH
func (c *Client) Exec(identity string, cmd string) (output []byte, err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		session, err := c.runtime[identity].NewSession()
		if err != nil {
			return
		}
		defer session.Close()
		output, err = session.Output(cmd)
	}()
	wg.Wait()
	return
}

// Get ssh client information
func (c *Client) GetConnectOption(identity string) (option *common.ConnectOption, err error) {
	if err := c.empty(identity); err != nil {
		return
	}
	option = c.options[identity]
	return
}

func (c *Client) GetRuntime(identity string) (client *ssh.Client, err error) {
	if err := c.empty(identity); err != nil {
		return
	}
	client = c.runtime[identity]
	return
}

func (c *Client) GetTunnelOption(identity string) (option []common.TunnelOption, err error) {
	if err := c.empty(identity); err != nil {
		return
	}
	option = *c.tunnels[identity]
	return
}

func (c *Client) All() []string {
	var keys []string
	for key := range c.options {
		keys = append(keys, key)
	}
	return keys
}

// Delete ssh client
func (c *Client) Delete(identity string) (err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	if c.tunnels[identity] != nil {
		c.closeTunnel(identity)
	}
	if c.runtime[identity] != nil {
		c.runtime[identity].Close()
	}
	delete(c.runtime, identity)
	delete(c.options, identity)
	return
}

// Tunnel setting
func (c *Client) SetTunnels(identity string, options []common.TunnelOption) (err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	if c.tunnels[identity] != nil {
		c.closeTunnel(identity)
	}
	c.tunnels[identity] = &options
	for _, tunnel := range options {
		go c.mutilTunnel(identity, tunnel)
	}
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
	localAddr := common.GetAddr(option.DstIp, uint(option.DstPort))
	remoteAddr := common.GetAddr(option.SrcIp, uint(option.SrcPort))
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
