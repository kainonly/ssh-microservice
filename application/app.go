package application

import (
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	_ "net/http/pprof"
	pb "ssh-microservice/api"
	"ssh-microservice/application/common"
	"ssh-microservice/application/controller"
)

func Application(dep common.Dependency) (err error) {
	cfg := dep.Config
	if cfg.Debug != "" {
		go http.ListenAndServe(cfg.Debug, nil)
	}
	var listen net.Listener
	if listen, err = net.Listen("tcp", cfg.Listen); err != nil {
		return
	}
	var logger *zap.Logger
	if logger, err = zap.NewProduction(); err != nil {
		return
	}
	defer logger.Sync()
	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpcZap.StreamServerInterceptor(logger),
		),
		grpc.UnaryInterceptor(
			grpcZap.UnaryServerInterceptor(logger),
		),
	)
	pb.RegisterAPIServer(
		server,
		controller.New(&dep),
	)
	go server.Serve(listen)
	return
}
