package client

import (
	"net"
	"sync"
)

type safeMapListener struct {
	sync.RWMutex
	Map map[string]map[string]*net.Listener
}

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
