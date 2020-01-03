package client

import (
	"golang.org/x/crypto/ssh"
	"ssh-gRPC/common"
)

// Test ssh client connection
func (c *Client) Testing(option common.ConnectOption) (sshClient *ssh.Client, err error) {
	return c.connect(option)
}
