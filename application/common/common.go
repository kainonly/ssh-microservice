package common

import (
	"go.uber.org/fx"
	"ssh-microservice/application/service/client"
	"ssh-microservice/application/service/schema"
	"ssh-microservice/config"
)

type Dependency struct {
	fx.In

	Config *config.Config
	Schema *schema.Schema
	Client *client.Client
}
