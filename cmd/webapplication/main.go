package main

import (
	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/di/webapplication"
	"go.uber.org/fx"
)

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		panic("Can't read configuration file: " + err.Error())
	}

	app := NewWebApplication(config)
	app.Run()
}

func NewWebApplication(config *configs.Config) *fx.App {
	app := fx.New(
		webapplication.WebApplicationModule{}.BuildOptions(config),
	)

	return app
}
