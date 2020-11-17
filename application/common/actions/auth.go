package actions

import (
	"encoding/base64"
	"golang.org/x/crypto/ssh"
	"ssh-microservice/config/options"
)

// Verify authentication method
func Auth(option options.ClientOption) (auth []ssh.AuthMethod, err error) {
	// Key method
	if option.PrivateKey != "" {
		var key []byte
		if key, err = base64.StdEncoding.DecodeString(option.PrivateKey); err != nil {
			return
		}
		var signer ssh.Signer
		if option.Passphrase == "" {
			if signer, err = ssh.ParsePrivateKey(key); err != nil {
				return
			}
		} else {
			var phrase []byte
			if phrase, err = base64.StdEncoding.DecodeString(option.Passphrase); err != nil {
				return
			}
			if signer, err = ssh.ParsePrivateKeyWithPassphrase(
				key,
				phrase,
			); err != nil {
				return
			}
		}
		auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
		return
	}
	// Password method
	if option.Password != "" {
		auth = []ssh.AuthMethod{
			ssh.Password(option.Password),
		}
	}
	return
}
