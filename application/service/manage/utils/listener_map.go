package utils

import (
	"net"
	"sync"
)

type ListenerMap struct {
	sync.RWMutex
	hashMap map[string]map[string]*net.Listener
}

func NewListenerMap() *ListenerMap {
	c := new(ListenerMap)
	c.hashMap = make(map[string]map[string]*net.Listener)
	return c
}

func (c *ListenerMap) Clear(identity string) {
	delete(c.hashMap, identity)
}

func (c *ListenerMap) Get(identity string, addr string) *net.Listener {
	c.RLock()
	listener := c.hashMap[identity][addr]
	c.RUnlock()
	return listener
}

func (c *ListenerMap) Set(identity string, addr string, listener *net.Listener) {
	c.Lock()
	if c.hashMap[identity] == nil {
		c.hashMap[identity] = make(map[string]*net.Listener)
	}
	c.hashMap[identity][addr] = listener
	c.Unlock()
}
