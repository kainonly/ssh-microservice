package client

import (
	"golang.org/x/crypto/ssh"
	"ssh-microservice/common"
)

// Test ssh client connection
func (c *Client) Testing(option common.ConnectOption) (sshClient *ssh.Client, err error) {
	return c.connect(option)
}
