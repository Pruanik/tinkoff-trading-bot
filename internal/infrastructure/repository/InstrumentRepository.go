package repository

import (
	"context"
	"errors"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database/mapping"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
	"gorm.io/gorm"
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

func (ir *InstrumentRepository) GetInstrumentByFigi(ctx context.Context, instrumentFigi string) (*model.Instrument, error) {
	var instrument model.Instrument
	res := ir.db.GetConnection().Model(&mapping.Instrument{}).Where("figi = ?", instrumentFigi).Take(&instrument)

	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			ir.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		}

		return nil, res.Error
	}

	return &instrument, nil
}
