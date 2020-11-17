package main

import (
	"go.uber.org/fx"
	"ssh-microservice/application"
	"ssh-microservice/bootstrap"
)

func main() {
	fx.New(
		//fx.NopLogger,
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeSchema,
			bootstrap.InitializeClient,
		),
		fx.Invoke(application.Application),
	).Run()
}
