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

func (c *ListenerMap) Put(identity string, addr string, listener *net.Listener) {
	c.Lock()
	if c.hashMap[identity] == nil {
		c.hashMap[identity] = make(map[string]*net.Listener)
	}
	c.hashMap[identity][addr] = listener
	c.Unlock()
}

func (c *ListenerMap) Lists(identity string) map[string]*net.Listener {
	c.RLock()
	lists := c.hashMap[identity]
	c.RUnlock()
	return lists
}

func (c *ListenerMap) Get(identity string, addr string) *net.Listener {
	c.RLock()
	listener := c.hashMap[identity][addr]
	c.RUnlock()
	return listener
}

func (c *ListenerMap) Remove(identity string) {
	delete(c.hashMap, identity)
}
