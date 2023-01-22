package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database/mapping"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

func NewLogRepository(db database.DatabaseInterface, logger log.LoggerInterface) repository.LogRepositoryInterface {
	return &LogRepository{db: db, logger: logger}
}

type LogRepository struct {
	db     database.DatabaseInterface
	logger log.LoggerInterface
}

func (lr *LogRepository) GetLogsDesc(ctx context.Context, lastLogId int, limit int) ([]model.Log, error) {
	var logs []model.Log

	res := lr.db.GetConnection().Model(&mapping.Log{}).Where("id > ?", lastLogId).Limit(limit).Order("id desc").Find(&logs)
	if res.Error != nil {
		lr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return logs, nil
}
