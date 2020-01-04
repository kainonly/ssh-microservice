package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"ssh-gRPC/client"
	"ssh-gRPC/controller"
	pb "ssh-gRPC/router"
)

func main() {
	listen, err := net.Listen("tcp", ":6060")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterRouterServer(
		server,
		controller.New(client.InjectClient()),
	)
	server.Serve(listen)
}
