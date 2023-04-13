package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type InstrumentSectorRepositoryInterface interface {
	Save(ctx context.Context, instrumentSector *model.InstrumentSector) (*model.InstrumentSector, error)

	GetInstrumentSectorByCode(ctx context.Context, code string) (*model.InstrumentSector, error)

	GetSectors(ctx context.Context) ([]model.InstrumentSector, error)
}
