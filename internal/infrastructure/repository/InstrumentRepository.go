package repository

import (
	"context"
	"errors"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
)

func NewInstrumentRepository(db database.DatabaseInterface) repository.InstrumentRepositoryInterface {
	return &InstrumentRepository{db: db}
}

type InstrumentRepository struct {
	db database.DatabaseInterface
}

func (ir *InstrumentRepository) Save(ctx context.Context, instrument *model.Instrument) (*model.Instrument, error) {
	res := ir.db.GetConnection().Save(instrument)
	if res.Error != nil {
		return nil, res.Error
	}

	return instrument, nil
}

func (ir *InstrumentRepository) GetInstruments(ctx context.Context) ([]model.Instrument, error) {
	var instruments []model.Instrument

	result := ir.db.GetConnection().Model(&model.Instrument{}).Order("type, name").Find(&instruments)
	if result.Error != nil {
		return instruments, errors.New("getInstruments: " + result.Error.Error())
	}

	return instruments, nil
}

func (ir *InstrumentRepository) GetInstrumentsByType(ctx context.Context, instrumentType string) ([]model.Instrument, error) {
	var instruments []model.Instrument

	result := ir.db.GetConnection().Model(&model.Instrument{}).Where("type = ?", instrumentType).Order("type, name").Find(&instruments)
	if result.Error != nil {
		return instruments, errors.New("getInstrumentsByType: " + result.Error.Error())
	}

	return instruments, nil
}
