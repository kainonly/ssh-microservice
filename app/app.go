package app

import (
	"google.golang.org/grpc"
	"net"
	"net/http"
	_ "net/http/pprof"
	"ssh-microservice/app/controller"
	"ssh-microservice/app/types"
	pb "ssh-microservice/router"
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

func (app *App) Start() (err error) {
	// Turn on debugging
	if app.option.Debug {
		go func() {
			http.ListenAndServe(":6060", nil)
		}()
	}
	// Start microservice
	listen, err := net.Listen("tcp", app.option.Listen)
	if err != nil {
		return
	}
	server := grpc.NewServer()
	pb.RegisterRouterServer(
		server,
		controller.New(),
	)
	server.Serve(listen)
	return
}
