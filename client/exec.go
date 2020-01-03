package client

import (
	"errors"
	"sync"
)

// Remotely execute commands via SSH
func (c *Client) Exec(identity string, cmd string) (output []byte, err error) {
	if c.options[identity] == nil || c.runtime[identity] == nil {
		err = errors.New("this identity does not exists")
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
