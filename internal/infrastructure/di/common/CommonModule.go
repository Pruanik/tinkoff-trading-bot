package common

import (
	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/service"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/repository"
	"go.uber.org/fx"
)

type CommonModule struct{}

func (cm CommonModule) BuildOptions(config *configs.Config) fx.Option {
	options := fx.Options(
		fx.Provide(
			func() *configs.Config {
				return config
			},
			database.NewDatabase,
			logger.NewLogger,

			repository.NewInstrumentRepository,
			repository.NewInstrumentSettingRepository,
			repository.NewInstrumentSectorRepository,
			repository.NewCandleRepository,

			service.NewInstrumentSectorService,
		),
	)

	return options
}
