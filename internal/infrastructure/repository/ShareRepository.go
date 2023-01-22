package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database/mapping"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

func NewShareRepository(db database.DatabaseInterface, logger log.LoggerInterface) repository.ShareRepositoryInterface {
	return &ShareRepository{db: db, logger: logger}
}

type ShareRepository struct {
	db     database.DatabaseInterface
	logger log.LoggerInterface
}

func (sr *ShareRepository) Save(ctx context.Context, share *model.Share) (*model.Share, error) {
	mappedShare := mapping.Share(*share)
	res := sr.db.GetConnection().Save(&mappedShare)
	if res.Error != nil {
		sr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return share, nil
}

func (sr *ShareRepository) GetShareByFigi(ctx context.Context, figi string) (*model.Share, error) {
	var share model.Share

	res := sr.db.GetConnection().Model(&mapping.Share{}).Where("figi = ?", figi).Find(&share)
	if res.Error != nil {
		sr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return &share, nil
}
