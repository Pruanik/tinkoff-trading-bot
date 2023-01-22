package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database/mapping"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

func NewInstrumentRepository(db database.DatabaseInterface, logger log.LoggerInterface) repository.InstrumentRepositoryInterface {
	return &InstrumentRepository{db: db, logger: logger}
}

type InstrumentRepository struct {
	db     database.DatabaseInterface
	logger log.LoggerInterface
}

func (ir *InstrumentRepository) Save(ctx context.Context, instrument *model.Instrument) (*model.Instrument, error) {
	mappedInstrument := mapping.Instrument(*instrument)
	res := ir.db.GetConnection().Save(&mappedInstrument)
	if res.Error != nil {
		ir.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return instrument, nil
}

func (ir *InstrumentRepository) GetInstruments(ctx context.Context) ([]model.Instrument, error) {
	var instruments []model.Instrument

	res := ir.db.GetConnection().Model(&mapping.Instrument{}).Order("type, name").Find(&instruments)
	if res.Error != nil {
		ir.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return instruments, nil
}

func (ir *InstrumentRepository) GetInstrumentsByType(ctx context.Context, instrumentType string) ([]model.Instrument, error) {
	var instruments []model.Instrument

	res := ir.db.GetConnection().Model(&mapping.Instrument{}).Where("type = ?", instrumentType).Order("type, name").Find(&instruments)
	if res.Error != nil {
		ir.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return instruments, nil
}

func (ir *InstrumentRepository) AreInstrumentsExistByType(ctx context.Context, instrumentType string) (bool, error) {
	var instruments []model.Instrument
	err := ir.db.GetConnection().Model(&mapping.Instrument{}).Where("type = ?", instrumentType).Limit(1).Find(&instruments).Error
	if err != nil {
		ir.logger.Error(log.LogCategoryDatabase, err.Error(), make(map[string]interface{}))
		return false, err
	}

	return len(instruments) > 0, nil
}
