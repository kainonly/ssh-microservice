package utils

import (
	"net"
	"sync"
)

type SyncConn struct {
	sync.RWMutex
	Map map[string]map[string]*net.Conn
}

func NewSyncConn() *SyncConn {
	listener := new(SyncConn)
	listener.Map = make(map[string]map[string]*net.Conn)
	return listener
}

func (c *SyncConn) Clear(identity string) {
	delete(c.Map, identity)
}

func (c *SyncConn) Get(identity string, addr string) *net.Conn {
	c.RLock()
	conn := c.Map[identity][addr]
	c.RUnlock()
	return conn
}

func (c *SyncConn) Set(identity string, addr string, conn *net.Conn) {
	c.Lock()
	if c.Map[identity] == nil {
		c.Map[identity] = make(map[string]*net.Conn)
	}
	c.Map[identity][addr] = conn
	c.Unlock()
}
