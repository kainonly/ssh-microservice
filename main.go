package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"ssh-microservice/client"
	"ssh-microservice/controller"
	pb "ssh-microservice/router"
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
