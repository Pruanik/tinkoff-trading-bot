package tinkoffinvest

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/module/tinkoffinvestconnection/tinkoffinvest"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
)

func NewInstrumentService(client investapi.InstrumentsServiceClient) tinkoffinvest.InstrumentServiceInterface {
	return &InstrumentService{client: client}
}

type InstrumentService struct {
	client investapi.InstrumentsServiceClient
}

func (is InstrumentService) GetBaseShares(ctx context.Context) (*investapi.SharesResponse, error) {
	instrumentRequest := investapi.InstrumentsRequest{InstrumentStatus: investapi.InstrumentStatus_INSTRUMENT_STATUS_BASE}
	shares, err := is.client.Shares(ctx, &instrumentRequest)
	if err != nil {
		return nil, err
	}

	return shares, nil
}

func (is InstrumentService) GetBaseCurrencies(ctx context.Context) (*investapi.CurrenciesResponse, error) {
	instrumentRequest := investapi.InstrumentsRequest{InstrumentStatus: investapi.InstrumentStatus_INSTRUMENT_STATUS_BASE}
	currencies, err := is.client.Currencies(ctx, &instrumentRequest)
	if err != nil {
		return nil, err
	}

	return currencies, nil
}
