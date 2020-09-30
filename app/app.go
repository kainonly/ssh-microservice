package app

import (
	"google.golang.org/grpc"
	"net"
	"net/http"
	_ "net/http/pprof"
	"ssh-microservice/app/controller"
	"ssh-microservice/app/manage"
	"ssh-microservice/app/types"
	pb "ssh-microservice/router"
)

func Application(option *types.Config) (err error) {
	// Turn on debugging
	if option.Debug != "" {
		go func() {
			http.ListenAndServe(option.Debug, nil)
		}()
	}
	// Start microservice
	listen, err := net.Listen("tcp", option.Listen)
	if err != nil {
		return
	}
	server := grpc.NewServer()
	manager, err := manage.NewClientManager()
	if err != nil {
		return
	}
	pb.RegisterRouterServer(
		server,
		controller.New(manager),
	)
	server.Serve(listen)
	return
}
