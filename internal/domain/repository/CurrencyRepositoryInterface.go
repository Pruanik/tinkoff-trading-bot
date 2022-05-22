package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type CurrencyRepositoryInterface interface {
	Save(ctx context.Context, currency *model.Currency) (*model.Currency, error)

	GetCurrencyByFigi(ctx context.Context, figi string) (*model.Currency, error)
}
