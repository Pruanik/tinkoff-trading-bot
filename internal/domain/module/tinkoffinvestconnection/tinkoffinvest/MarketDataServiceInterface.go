package tinkoffinvest

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/grpc/investapi"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MarketDataServiceInterface interface {
	GetHistoricalCandlesByFigi(ctx context.Context, figi string, from *timestamppb.Timestamp) (*investapi.GetCandlesResponse, error)
}
