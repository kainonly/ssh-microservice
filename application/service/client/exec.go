package client

import "golang.org/x/crypto/ssh"

func (c *Client) Exec(identity string, cmd string) (output []byte, err error) {
	if c.Options.Empty(identity) {
		return nil, NotExists
	}
	var session *ssh.Session
	if session, err = c.Clients.Get(identity).NewSession(); err != nil {
		return
	}
	defer session.Close()
	if output, err = session.Output(cmd); err != nil {
		return
	}
	return
}
