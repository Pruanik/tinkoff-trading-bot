package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
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
	res := cr.db.GetConnection().Save(candle)
	if res.Error != nil {
		cr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return candle, nil
}

func (cr *CandleRepository) GetLastCandleByFigi(ctx context.Context, figi string) (*model.Candle, error) {
	var candle model.Candle

	res := cr.db.GetConnection().Model(&model.Candle{}).Where("figi = ?", figi).Order("id desc").Limit(1).Find(&candle)
	if res.Error != nil {
		cr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return &candle, nil
}
