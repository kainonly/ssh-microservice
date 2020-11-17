package actions

import (
	"golang.org/x/crypto/ssh"
	pb "ssh-microservice/api"
)

// Verify authentication method
func Auth(option *pb.Option) (auth []ssh.AuthMethod, err error) {
	// Priority detection key method
	if len(option.PrivateKey) != 0 {
		var signer ssh.Signer
		if len(option.Passphrase) == 0 {
			if signer, err = ssh.ParsePrivateKey(option.PrivateKey); err != nil {
				return
			}
		} else {
			if signer, err = ssh.ParsePrivateKeyWithPassphrase(
				option.PrivateKey,
				option.Passphrase,
			); err != nil {
				return
			}
		}
		auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
		return
	}
	// Use password
	if option.Password != "" {
		auth = []ssh.AuthMethod{
			ssh.Password(option.Password),
		}
	}
	return
}
