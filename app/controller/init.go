package controller

import (
	"ssh-microservice/app/manage"
	pb "ssh-microservice/router"
)

type controller struct {
	pb.UnimplementedRouterServer
	m *manage.ClientManager
}

func New(manager *manage.ClientManager) *controller {
	c := new(controller)
	c.m = manager
	return c
}
