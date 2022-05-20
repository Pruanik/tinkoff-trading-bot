package webapplication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Pruanik/tinkoff-trading-bot/api/server"
	apiRouter "github.com/Pruanik/tinkoff-trading-bot/api/server/api"
	webRouter "github.com/Pruanik/tinkoff-trading-bot/api/server/web"
	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/di/common"
	apiHandler "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/httphandler/api"
	webHandler "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/httphandler/web"
	"go.uber.org/fx"
)

type WebApplicationModule struct{}

func (wam WebApplicationModule) BuildOptions(config *configs.Config) fx.Option {
	options := fx.Options(
		common.CommonModule{}.BuildOptions(config),
		fx.Provide(
			server.NewRouter,
			server.NewServer,
			webRouter.NewWebRouter,
			apiRouter.NewApiRouter,
			webHandler.NewHomeHandler,
			apiHandler.NewLogApiHandler,
		),
		fx.Invoke(
			server.RegisterRoutes,
			wam.startApplicationServer,
		),
	)

	return options
}

func (wam WebApplicationModule) startApplicationServer(lc fx.Lifecycle, server *http.Server) {
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
