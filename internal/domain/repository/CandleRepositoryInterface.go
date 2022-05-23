package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type CandleRepositoryInterface interface {
	Save(ctx context.Context, candle *model.Candle) (*model.Candle, error)

	GetLastCandleByFigi(ctx context.Context, figi string) (*model.Candle, error)
}
