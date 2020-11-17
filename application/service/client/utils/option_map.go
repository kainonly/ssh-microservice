package utils

import (
	"ssh-microservice/config/options"
)

type OptionMap struct {
	hashMap map[string]*options.ClientOption
}

func NewOptionMap() *OptionMap {
	c := new(OptionMap)
	c.hashMap = make(map[string]*options.ClientOption)
	return c
}

func (c *OptionMap) Put(identity string, option *options.ClientOption) {
	c.hashMap[identity] = option
}

func (c *OptionMap) SetTunnels(identity string, tunnels []options.TunnelOption) {
	c.hashMap[identity].Tunnels = tunnels
}

func (c *OptionMap) Lists() map[string]*options.ClientOption {
	return c.hashMap
}

func (c *OptionMap) Empty(identity string) bool {
	return c.hashMap[identity] == nil
}

func (c *OptionMap) TunnelsIsEmpty(identity string) bool {
	return len(c.hashMap[identity].Tunnels) == 0
}

func (c *OptionMap) Get(identity string) *options.ClientOption {
	return c.hashMap[identity]
}

func (c *OptionMap) Remove(identity string) {
	delete(c.hashMap, identity)
}
