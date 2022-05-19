package tinkoffinvestconnection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/di/common"
	"go.uber.org/fx"
)

type TinkoffInvestConnectionModule struct{}

func (tic TinkoffInvestConnectionModule) BuildOptions(config *configs.Config) fx.Option {
	options := fx.Options(
		common.CommonModule{}.BuildOptions(config),
		fx.Provide(),
		fx.Invoke(),
	)

	return options
}

func (tic TinkoffInvestConnectionModule) startApplicationServer(lc fx.Lifecycle, server *http.Server) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					err := server.ListenAndServe()
					if err != nil {
						fmt.Printf("HTTP Server has error while it try to start! Error: %s", err)
					}
				}()
				return nil
			},
		},
	)
}
