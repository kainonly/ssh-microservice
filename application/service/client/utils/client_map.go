package utils

import (
	"golang.org/x/crypto/ssh"
)

type ClientMap struct {
	hashMap map[string]*ssh.Client
}

func NewClientMap() *ClientMap {
	c := new(ClientMap)
	c.hashMap = make(map[string]*ssh.Client)
	return c
}

func (c *ClientMap) Put(identity string, value *ssh.Client) {
	c.hashMap[identity] = value
}

func (c *ClientMap) Empty(identity string) bool {
	return c.hashMap[identity] == nil
}

func (c *ClientMap) Get(identity string) *ssh.Client {
	return c.hashMap[identity]
}

func (c *ClientMap) Remove(identity string) {
	delete(c.hashMap, identity)
}
