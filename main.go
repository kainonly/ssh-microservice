package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"ssh-microservice/client"
	"ssh-microservice/common"
	"ssh-microservice/controller"
	pb "ssh-microservice/router"
)

func main() {
	server := grpc.NewServer()
	common.InitLevelDB("data")
	common.InitBufPool()
	pb.RegisterRouterServer(
		server,
		controller.New(client.InjectClient()),
	)
	listen, err := net.Listen("tcp", ":6060")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server.Serve(listen)
}
