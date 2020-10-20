package manage

func (c *ClientManager) Delete(identity string) (err error) {
	if err = c.empty(identity); err != nil {
		return
	}
	if c.tunnels[identity] != nil {
		c.closeTunnel(identity)
	}
	if c.runtime[identity] != nil {
		c.runtime[identity].Close()
	}
	if c.keepalive[identity] != nil {
		c.keepalive[identity].Stop()
	}
	delete(c.runtime, identity)
	delete(c.options, identity)
	delete(c.keepalive, identity)
	return c.schema.Delete(identity)
}
