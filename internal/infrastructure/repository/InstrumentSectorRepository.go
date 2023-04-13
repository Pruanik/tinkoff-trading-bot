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

func NewInstrumentSectorRepository(db database.DatabaseInterface, logger log.LoggerInterface) repository.InstrumentSectorRepositoryInterface {
	return &InstrumentSectorRepository{db: db, logger: logger}
}

type InstrumentSectorRepository struct {
	db     database.DatabaseInterface
	logger log.LoggerInterface
}

func (isr *InstrumentSectorRepository) Save(ctx context.Context, instrumentSector *model.InstrumentSector) (*model.InstrumentSector, error) {
	mappedInstrumentSector := mapping.InstrumentSector(*instrumentSector)
	res := isr.db.GetConnection().Save(&mappedInstrumentSector)
	if res.Error != nil {
		isr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}
	// mapping revert?
	return instrumentSector, nil
}

func (isr *InstrumentSectorRepository) GetInstrumentSectorByCode(ctx context.Context, code string) (*model.InstrumentSector, error) {
	var instrumentSector model.InstrumentSector
	res := isr.db.GetConnection().Model(&mapping.InstrumentSector{}).Where("code = ?", code).Take(&instrumentSector)

	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			isr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		}

		return nil, res.Error
	}

	return &instrumentSector, nil
}

func (isr *InstrumentSectorRepository) GetSectors(ctx context.Context) ([]model.InstrumentSector, error) {
	var sectors []model.InstrumentSector

	res := isr.db.GetConnection().Model(&mapping.InstrumentSector{}).Order("name").Find(&sectors)
	if res.Error != nil {
		isr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return sectors, nil
}
