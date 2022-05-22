package tinkoffinvest

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
)

type InstrumentServiceInterface interface {
	GetBaseShares(ctx context.Context) (*investapi.SharesResponse, error)

	GetBaseCurrencies(ctx context.Context) (*investapi.CurrenciesResponse, error)
}
