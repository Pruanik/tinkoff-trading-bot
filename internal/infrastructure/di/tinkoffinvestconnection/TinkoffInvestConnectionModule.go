package tinkoffinvestconnection

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/fillinghistoricaldata/candles"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/fillinginstrumentsinfo"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/fillinginstrumentsinfo/fillingcurrenciesinfo"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/fillinginstrumentsinfo/fillingsharesinfo"
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
			repository.NewInstrumentRepository,
			repository.NewShareRepository,
			repository.NewCurrencyRepository,
			fillingcurrenciesinfo.NewFillingCurrenciesInfo,
			fillingsharesinfo.NewFillingSharesInfo,
			fillinginstrumentsinfo.NewFillingInstrumentsInfo,
			candles.NewFillingHistoricalCandlesData,
		),
		fx.Invoke(
			tic.startControllSettingsUpdate,
			tic.startCollectMarketdata,
		),
	)

	return options
}

func (tic TinkoffInvestConnectionModule) startControllSettingsUpdate(
	lc fx.Lifecycle,
	fillingInstrumentsInfo fillinginstrumentsinfo.FillingInstrumentsInfoInterface,
	fillingHistoricalCandlesData candles.FillingHistoricalCandlesDataInterface,
) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					fillingInstrumentsInfo.LoadInfo(ctx)
					fillingHistoricalCandlesData.FillingHistoricalData(ctx)
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
