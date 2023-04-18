package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type InstrumentRepositoryInterface interface {
	Save(ctx context.Context, instrument *model.Instrument) (*model.Instrument, error)

	GetInstruments(ctx context.Context) ([]model.Instrument, error)

	GetInstrumentsByType(ctx context.Context, instrumentType string) ([]model.Instrument, error)

	GetInstrumentsBySectorId(ctx context.Context, sectorId string) ([]model.Instrument, error)

	GetInstrumentsBySectorIdAndType(ctx context.Context, sectorId string, instrumentType string) ([]model.Instrument, error)

	GetInstrumentByFigi(ctx context.Context, instrumentFigi string) (*model.Instrument, error)

	GetInstrumentTypes(ctx context.Context) ([]string, error)

	GetInstrumentTypesBySectorId(ctx context.Context, sectorId string) ([]string, error)
}
