package utils

import (
	"net"
	"sync"
)

type SyncListener struct {
	sync.RWMutex
	Map map[string]map[string]*net.Listener
}

func NewSyncListener() *SyncListener {
	c := new(SyncListener)
	c.Map = make(map[string]map[string]*net.Listener)
	return c
}

func (c *SyncListener) Clear(identity string) {
	delete(c.Map, identity)
}

func (c *SyncListener) Get(identity string, addr string) *net.Listener {
	c.RLock()
	listener := c.Map[identity][addr]
	c.RUnlock()
	return listener
}

func (c *SyncListener) Set(identity string, addr string, listener *net.Listener) {
	c.Lock()
	if c.Map[identity] == nil {
		c.Map[identity] = make(map[string]*net.Listener)
	}
	c.Map[identity][addr] = listener
	c.Unlock()
}
