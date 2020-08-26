package controller

import (
	pb "ssh-microservice/router"
)

type controller struct {
	pb.UnimplementedRouterServer
}

func New() *controller {
	c := new(controller)
	return c
}
