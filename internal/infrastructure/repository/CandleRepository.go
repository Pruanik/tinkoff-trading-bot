package repository

import (
	"context"
	"time"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database/mapping"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

func NewCandleRepository(db database.DatabaseInterface, logger log.LoggerInterface) repository.CandleRepositoryInterface {
	return &CandleRepository{db: db, logger: logger}
}

type CandleRepository struct {
	db     database.DatabaseInterface
	logger log.LoggerInterface
}

func (cr *CandleRepository) Save(ctx context.Context, candle *model.Candle) (*model.Candle, error) {
	mappedCandle := mapping.Candle(*candle)
	res := cr.db.GetConnection().Save(&mappedCandle)
	if res.Error != nil {
		cr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return candle, nil
}

func (cr *CandleRepository) GetLastCandleByFigi(ctx context.Context, figi string) (*model.Candle, error) {
	var candle model.Candle

	res := cr.db.GetConnection().Model(&mapping.Candle{}).Where("figi = ?", figi).Order("id desc").Limit(1).Find(&candle)
	if res.Error != nil {
		cr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return &candle, nil
}

func (cr *CandleRepository) GetCandlesByFigiFromTime(ctx context.Context, figi string, time time.Time) ([]model.Candle, error) {
	var candles []model.Candle

	res := cr.db.GetConnection().Model(&mapping.Candle{}).Where("figi = ? and timestamp > ?", figi, time.Format("2006-01-02 15:04:05")).Order("timestamp").Find(&candles)
	if res.Error != nil {
		cr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return candles, nil
}

func (cr *CandleRepository) GetCandlesByFigiFromLastId(ctx context.Context, figi string, lastId int) ([]model.Candle, error) {
	var candles []model.Candle

	res := cr.db.GetConnection().Model(&mapping.Candle{}).Where("figi = ? and id > ?", figi, lastId).Order("timestamp").Find(&candles)
	if res.Error != nil {
		cr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return candles, nil
}
