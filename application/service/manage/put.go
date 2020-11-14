package manage

import (
	"encoding/base64"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"ssh-microservice/app/actions"
	"ssh-microservice/app/types"
)

func (c *ClientManager) Put(identity string, option types.SshOption) (err error) {
	if c.options[identity] != nil && c.runtime[identity] != nil {
		if err = c.Delete(identity); err != nil {
			return
		}
	}
	c.options[identity] = &option
	c.runtime[identity], err = actions.Connect(option)
	if err != nil {
		return
	}
	c.keepalive[identity] = cron.New(cron.WithSeconds())
	c.keepalive[identity].AddFunc("*/30 * * * * *", func() {
		var session *ssh.Session
		session, err = c.runtime[identity].NewSession()
		if err != nil {
			logrus.Error("the [", identity, "] keeplive failed")
		}
		defer session.Close()
		_, err = session.Output("uptime")
		if err != nil {
			logrus.Error("the [", identity, "] keeplive failed")
		}
	})
	c.keepalive[identity].Start()
	return c.schema.Update(types.ClientOption{
		Identity:   identity,
		Host:       option.Host,
		Port:       option.Port,
		Username:   option.Username,
		Password:   option.Password,
		Key:        base64.StdEncoding.EncodeToString(option.Key),
		PassPhrase: base64.StdEncoding.EncodeToString(option.PassPhrase),
		Tunnels:    nil,
	})
}
