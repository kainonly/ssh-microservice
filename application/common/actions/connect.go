package actions

import (
	"golang.org/x/crypto/ssh"
	"ssh-microservice/config/options"
)

// Connect to server
func Connect(option options.ClientOption) (client *ssh.Client, err error) {
	auth, err := Auth(option)
	if err != nil {
		return
	}
	config := ssh.ClientConfig{
		User:            option.Username,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := GetAddr(option.Host, uint(option.Port))
	client, err = ssh.Dial("tcp", addr, &config)
	return
}
