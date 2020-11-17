package client

func (c *Client) Delete(identity string) (err error) {
	if c.Options.Empty(identity) {
		return
	}
	if !c.Options.TunnelsIsEmpty(identity) {
		c.closeTunnel(identity)
	}
	if !c.Clients.Empty(identity) {
		c.Clients.Get(identity).Close()
	}
	if !c.keepalive.Empty(identity) {
		c.keepalive.Get(identity).Stop()
	}
	c.Options.Remove(identity)
	c.Clients.Remove(identity)
	c.keepalive.Remove(identity)
	return c.schema.Delete(identity)
}
