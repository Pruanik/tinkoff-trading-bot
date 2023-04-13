package service

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type InstrumentSectorServiceInterface interface {
	Create(ctx context.Context, code string) (*model.InstrumentSector, error)

	CreateIfNotExist(ctx context.Context, code string) (*model.InstrumentSector, error)
}
