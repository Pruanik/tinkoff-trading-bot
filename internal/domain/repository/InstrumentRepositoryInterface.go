package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type InstrumentRepositoryInterface interface {
	Save(ctx context.Context, instrument *model.Instrument) (*model.Instrument, error)

	GetInstruments(ctx context.Context) ([]model.Instrument, error)

	GetInstrumentsByType(ctx context.Context, instrumentType string) ([]model.Instrument, error)

	GetInstrumentByFigi(ctx context.Context, instrumentFigi string) (*model.Instrument, error)
}
