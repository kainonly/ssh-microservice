package app

import (
	"net"
	"ssh-microservice/app/types"
	"sync"
)

type App struct {
	option  *types.Config
	bufPool *sync.Pool
}

func New(config types.Config) *App {
	app := new(App)
	app.option = &config
	app.bufPool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, config.Pool*1024)
		},
	}
	return app
}

func (c *App) Start() (err error) {
	_, err = net.Listen("tcp", c.option.Listen)
	if err != nil {
		return
	}
	return
}
