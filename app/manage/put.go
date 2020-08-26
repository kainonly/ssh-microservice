package manage

import (
	"ssh-microservice/app/actions"
	"ssh-microservice/app/types"
	"sync"
)

func (c *ClientManager) Put(identity string, option types.SshOption) (err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.options[identity] = &option
		c.runtime[identity], err = actions.Connect(option)
		if err != nil {
			return
		}
		go c.runtime[identity].Wait()
	}()
	wg.Wait()
	return
}
