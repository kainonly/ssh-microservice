package common

import (
	"go.uber.org/fx"
	"ssh-microservice/config"
)

type Dependency struct {
	fx.In

	Config *config.Config
}
