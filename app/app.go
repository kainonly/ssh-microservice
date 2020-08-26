package app

import (
	"ssh-microservice/app/types"
	"sync"
)

type App struct {
	bufPool *sync.Pool
}

func New(config types.Config) *App {
	app := new(App)
	app.bufPool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, config.Pool*1024)
		},
	}
	return app
}

func (c *App) Start() {

}
