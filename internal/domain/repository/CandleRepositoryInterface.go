package repository

import (
	"context"
	"time"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type CandleRepositoryInterface interface {
	Save(ctx context.Context, candle *model.Candle) (*model.Candle, error)

	GetLastCandleByFigi(ctx context.Context, figi string) (*model.Candle, error)

	GetCandlesByFigiFromTime(ctx context.Context, figi string, time time.Time) ([]model.Candle, error)

	GetCandlesByFigiFromLastId(ctx context.Context, figi string, lastId int) ([]model.Candle, error)
}
