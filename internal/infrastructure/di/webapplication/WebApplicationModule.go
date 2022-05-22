package webapplication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Pruanik/tinkoff-trading-bot/api/server"
	apiRouter "github.com/Pruanik/tinkoff-trading-bot/api/server/api"
	webRouter "github.com/Pruanik/tinkoff-trading-bot/api/server/web"
	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/di/common"
	apiHandler "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/httphandler/api"
	webHandler "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/httphandler/web"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/repository"
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
			webHandler.NewPageNotFoundHandler,
			apiHandler.NewLogApiHandler,
			apiHandler.NewInstrumentApiHandler,
			repository.NewInstrumentRepository,
			repository.NewLogRepository,
			builder.NewHttpResponseBuilder,
			builder.NewGetInstrumentsBodyBuilder,
			builder.NewGetLogsBodyBuilder,
		),
		fx.Invoke(
			server.RegisterRoutes,
			wam.startApplicationServer,
		),
	)

	return options
}

func (wam WebApplicationModule) startApplicationServer(lc fx.Lifecycle, server *http.Server, logger log.LoggerInterface) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					logger.Info(log.LogCategorySystem, "Start HTTP Server", make(map[string]interface{}))
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
