package client

import (
	"github.com/robfig/cron/v3"
	"golang.org/x/crypto/ssh"
	"log"
	"ssh-microservice/application/common/actions"
	"ssh-microservice/config/options"
)

func (c *Client) Put(option options.ClientOption) (err error) {
	identity := option.Identity
	if !c.Options.Empty(identity) {
		if err = c.Delete(identity); err != nil {
			return
		}
	}
	c.Options.Put(identity, &option)
	var client *ssh.Client
	if client, err = actions.Connect(option); err != nil {
		return
	}
	c.Clients.Put(identity, client)
	schedule := cron.New(cron.WithSeconds())
	schedule.AddFunc("*/30 * * * * *", func() {
		var session *ssh.Session
		if session, err = client.NewSession(); err != nil {
			log.Println("the [", identity, "] keeplive failed")
		}
		defer session.Close()
		if _, err = session.Output("uptime"); err != nil {
			log.Println("the [", identity, "] keeplive failed")
		}
	})
	schedule.Start()
	c.keepalive.Put(identity, schedule)
	return c.schema.Update(option)
}
