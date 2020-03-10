package client

import (
	"net"
	"sync"
)

type safeMapConn struct {
	sync.RWMutex
	Map map[string]map[string]*net.Conn
}

func newSafeMapConn() *safeMapConn {
	listener := new(safeMapConn)
	listener.Map = make(map[string]map[string]*net.Conn)
	return listener
}

func (s *safeMapConn) Clear(identity string) {
	delete(s.Map, identity)
}

func (s *safeMapConn) Get(identity string, addr string) *net.Conn {
	s.RLock()
	conn := s.Map[identity][addr]
	s.RUnlock()
	return conn
}

func (s *safeMapConn) Set(identity string, addr string, conn *net.Conn) {
	s.Lock()
	if s.Map[identity] == nil {
		s.Map[identity] = make(map[string]*net.Conn)
	}
	s.Map[identity][addr] = conn
	s.Unlock()
}
