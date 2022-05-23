package tinkoffinvest

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

func NewInstrumentService(client investapi.InstrumentsServiceClient, logger log.LoggerInterface) tinkoffinvest.InstrumentServiceInterface {
	return &InstrumentService{client: client, logger: logger}
}

type InstrumentService struct {
	client investapi.InstrumentsServiceClient
	logger log.LoggerInterface
}

func (is InstrumentService) GetBaseShares(ctx context.Context) (*investapi.SharesResponse, error) {
	instrumentRequest := investapi.InstrumentsRequest{InstrumentStatus: investapi.InstrumentStatus_INSTRUMENT_STATUS_BASE}
	shares, err := is.client.Shares(ctx, &instrumentRequest)
	if err != nil {
		is.logger.Error(log.LogCategoryGrpcTinkoff, err.Error(), map[string]interface{}{"service": "InstrumentService", "method": "GetBaseShares"})
		return nil, err
	}

	return shares, nil
}

func (is InstrumentService) GetBaseCurrencies(ctx context.Context) (*investapi.CurrenciesResponse, error) {
	instrumentRequest := investapi.InstrumentsRequest{InstrumentStatus: investapi.InstrumentStatus_INSTRUMENT_STATUS_BASE}
	currencies, err := is.client.Currencies(ctx, &instrumentRequest)
	if err != nil {
		is.logger.Error(log.LogCategoryGrpcTinkoff, err.Error(), map[string]interface{}{"service": "InstrumentService", "method": "GetBaseCurrencies"})
		return nil, err
	}

	return currencies, nil
}
