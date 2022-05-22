package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type ShareRepositoryInterface interface {
	Save(ctx context.Context, share *model.Share) (*model.Share, error)

	GetShareByFigi(ctx context.Context, figi string) (*model.Share, error)
}
