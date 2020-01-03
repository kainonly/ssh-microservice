package client

import (
	"net"
	"sync"
)

type (
	safeMapListener struct {
		sync.RWMutex
		Map map[string]map[string]*net.Listener
	}
	safeMapConn struct {
		sync.RWMutex
		Map map[string]map[string]*net.Conn
	}
)

func newSafeMapListener() *safeMapListener {
	listener := new(safeMapListener)
	listener.Map = make(map[string]map[string]*net.Listener)
	return listener
}

func (s *safeMapListener) Clear(identity string) {
	s.Map[identity] = make(map[string]*net.Listener)
}

func (s *safeMapListener) Get(identity string, addr string) *net.Listener {
	s.RLock()
	listener := s.Map[identity][addr]
	s.RUnlock()
	return listener
}

func (s *safeMapListener) Set(identity string, addr string, listener *net.Listener) {
	s.Lock()
	s.Map[identity][addr] = listener
	s.Unlock()
}

func newSafeMapConn() *safeMapConn {
	listener := new(safeMapConn)
	listener.Map = make(map[string]map[string]*net.Conn)
	return listener
}

func (s *safeMapConn) Clear(identity string) {
	s.Map[identity] = make(map[string]*net.Conn)
}

func (s *safeMapConn) Get(identity string, addr string) *net.Conn {
	s.RLock()
	conn := s.Map[identity][addr]
	s.RUnlock()
	return conn
}

func (s *safeMapConn) Set(identity string, addr string, conn *net.Conn) {
	s.Lock()
	s.Map[identity][addr] = conn
	s.Unlock()
}
