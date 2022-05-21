package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type InstrumentRepositoryInterface interface {
	GetInstruments(ctx context.Context) ([]model.Instrument, error)

	GetInstrumentsByType(ctx context.Context, instrumentType string) ([]model.Instrument, error)
}
