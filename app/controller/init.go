package controller

import (
	"ssh-microservice/app/manage"
	pb "ssh-microservice/router"
)

type controller struct {
	pb.UnimplementedRouterServer
	m *manage.ClientManager
}

func New() *controller {
	c := new(controller)
	c.m = manage.NewClientManager()
	return c
}
