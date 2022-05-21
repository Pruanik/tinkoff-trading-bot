package tinkoffinvestconnection

import (
	"context"
	"fmt"

	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/di/common"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
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
		),
		fx.Invoke(
			tic.startCollectMarketdata,
		),
	)

	return options
}

func (tic TinkoffInvestConnectionModule) startCollectMarketdata(lc fx.Lifecycle, instrument investapi.InstrumentsServiceClient) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					instrumentRequest := investapi.InstrumentsRequest{InstrumentStatus: investapi.InstrumentStatus_INSTRUMENT_STATUS_BASE}
					shares, err := instrument.Shares(ctx, &instrumentRequest)
					if err != nil {
						fmt.Printf("HTTP Server has error while it try to start! Error: %s", err)
					}
					fmt.Println(shares)
				}()
				return nil
			},
		},
	)
}
