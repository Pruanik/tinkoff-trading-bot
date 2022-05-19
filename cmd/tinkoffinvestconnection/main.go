package main

import (
	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/di/tinkoffinvestconnection"
	"go.uber.org/fx"
)

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		panic("Can't read configuration file: " + err.Error())
	}

	app := NewTinkoffInvestConnection(config)
	app.Run()
}

func NewTinkoffInvestConnection(config *configs.Config) *fx.App {
	app := fx.New(
		tinkoffinvestconnection.TinkoffInvestConnectionModule{}.BuildOptions(config),
	)

	return app
}
