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

func NewInstrumentSettingRepository(db database.DatabaseInterface, logger log.LoggerInterface) repository.InstrumentSettingRepositoryInterface {
	return &InstrumentSettingRepository{db: db, logger: logger}
}

type InstrumentSettingRepository struct {
	db     database.DatabaseInterface
	logger log.LoggerInterface
}

func (isr *InstrumentSettingRepository) Save(ctx context.Context, instrumentSetting *model.InstrumentSetting) (*model.InstrumentSetting, error) {
	mappedInstrumentSetting := mapping.InstrumentSetting(*instrumentSetting)
	res := isr.db.GetConnection().Save(&mappedInstrumentSetting)
	if res.Error != nil {
		isr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return instrumentSetting, nil
}

func (isr *InstrumentSettingRepository) GetInstrumentSettingByFigi(ctx context.Context, figi string) (*model.InstrumentSetting, error) {
	var instrumentSetting model.InstrumentSetting
	res := isr.db.GetConnection().Model(&mapping.InstrumentSetting{}).Where("figi = ?", figi).Take(&instrumentSetting)

	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			isr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		}

		return nil, res.Error
	}

	return &instrumentSetting, nil
}

func (isr *InstrumentSettingRepository) Update(ctx context.Context, instrumentSetting *model.InstrumentSetting) (*model.InstrumentSetting, error) {
	instrumentSettingInDb, _ := isr.GetInstrumentSettingByFigi(ctx, instrumentSetting.Figi)

	if instrumentSettingInDb == nil {
		isr.Save(ctx, instrumentSetting)
	} else {
		isr.db.GetConnection().Model(&mapping.InstrumentSetting{}).Where("figi = ?", instrumentSetting.Figi).Update("is_data_collecting", instrumentSetting.IsDataCollecting)
	}

	return instrumentSetting, nil
}

func (isr *InstrumentSettingRepository) GetInstrumentsSettingsWhereIsCollectingTrue(ctx context.Context) ([]model.InstrumentSettingWithName, error) {
	var result []model.InstrumentSettingWithName
	res := isr.db.GetConnection().Model(&mapping.Instrument{}).Select("instruments.figi, instruments.name, instrument_settings.is_data_collecting, instrument_settings.created_at").Joins("left join instrument_settings on instrument_settings.figi = instruments.figi").Where("instrument_settings.is_data_collecting = ?", true).Scan(&result)
	if res.Error != nil {
		isr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return result, nil
}
