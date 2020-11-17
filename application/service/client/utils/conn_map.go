package utils

import (
	"net"
	"sync"
)

type ConnMap struct {
	sync.RWMutex
	hashMap map[string]map[string]*net.Conn
}

func NewConnMap() *ConnMap {
	listener := new(ConnMap)
	listener.hashMap = make(map[string]map[string]*net.Conn)
	return listener
}

func (c *ConnMap) Put(identity string, addr string, conn *net.Conn) {
	c.Lock()
	if c.hashMap[identity] == nil {
		c.hashMap[identity] = make(map[string]*net.Conn)
	}
	c.hashMap[identity][addr] = conn
	c.Unlock()
}

func (c *ConnMap) Lists(identity string) map[string]*net.Conn {
	c.RLock()
	lists := c.hashMap[identity]
	c.RUnlock()
	return lists
}

func (c *ConnMap) Get(identity string, addr string) *net.Conn {
	c.RLock()
	conn := c.hashMap[identity][addr]
	c.RUnlock()
	return conn
}

func (c *ConnMap) Remove(identity string) {
	delete(c.hashMap, identity)
}
