package client

import (
	"ssh-microservice/common"
	"sync"
)

// Add or modify the ssh client
func (c *Client) Put(identity string, option common.ConnectOption) (err error) {
	err = c.Delete(identity)
	if err != nil {
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
	err = common.SetTemporary(common.ConfigOption{
		Connect: c.options,
		Tunnel:  c.tunnels,
	})
	return
}
