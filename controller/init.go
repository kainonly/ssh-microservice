package controller

import (
	"ssh-microservice/client"
	pb "ssh-microservice/router"
)

type controller struct {
	pb.UnimplementedRouterServer
	client client.Client
}

func New(client *client.Client) *controller {
	c := new(controller)
	c.client = *client
	return c
}
