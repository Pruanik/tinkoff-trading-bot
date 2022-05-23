package tinkoffinvest

import (
	"context"
	"time"

	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
)

type MarketDataServiceInterface interface {
	GetHistoricalCandlesByFigi(ctx context.Context, figi string, from time.Time, to time.Time) (*investapi.GetCandlesResponse, error)
}
