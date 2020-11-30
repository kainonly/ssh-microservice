package application

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	_ "net/http/pprof"
	pb "ssh-microservice/api"
	"ssh-microservice/application/common"
	"ssh-microservice/application/controller"
	"ssh-microservice/bootstrap"
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
			grpc_middleware.ChainStreamServer(
				grpcZap.StreamServerInterceptor(logger),
				grpcRecovery.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpcZap.UnaryServerInterceptor(logger),
				grpcRecovery.UnaryServerInterceptor(),
			),
		),
	)
	pb.RegisterAPIServer(
		server,
		controller.New(&dep),
	)
	go server.Serve(listen)
	if cfg.Gateway != "" {
		go bootstrap.ApiGateway(cfg)
	}
	return
}
