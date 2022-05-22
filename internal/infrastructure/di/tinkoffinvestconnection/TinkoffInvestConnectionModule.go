package tinkoffinvestconnection

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/di/common"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/tinkoffinvest"
	"go.uber.org/fx"
)

type TinkoffInvestConnectionModule struct{}

func (tic TinkoffInvestConnectionModule) BuildOptions(config *configs.Config) fx.Option {
	options := fx.Options(
		common.CommonModule{}.BuildOptions(config),
		fx.Provide(
			grpc.NewGrpcConnection,
			investapi.NewInstrumentsServiceClient,
			investapi.NewMarketDataServiceClient,
			investapi.NewMarketDataStreamServiceClient,
			investapi.NewOperationsServiceClient,
			investapi.NewOrdersServiceClient,
			investapi.NewSandboxServiceClient,
			investapi.NewStopOrdersServiceClient,
			investapi.NewUsersServiceClient,
			tinkoffinvest.NewInstrumentService,
			tinkoffinvestconnection.NewCheckInstrumentsData,
			repository.NewInstrumentRepository,
		),
		fx.Invoke(
			tic.startControllSettingsUpdate,
			tic.startCollectMarketdata,
		),
	)

	return options
}

func (tic TinkoffInvestConnectionModule) startControllSettingsUpdate(lc fx.Lifecycle, checkInstrumentsData *tinkoffinvestconnection.CheckInstrumentsData) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					checkInstrumentsData.CheckDataExistAndLoad(ctx)
				}()
				return nil
			},
		},
	)
}

func (tic TinkoffInvestConnectionModule) startCollectMarketdata(lc fx.Lifecycle, instrument investapi.InstrumentsServiceClient) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {

				}()
				return nil
			},
		},
	)
}
