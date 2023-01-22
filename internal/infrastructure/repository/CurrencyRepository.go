package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database/mapping"
	log "github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/logger"
)

func NewCurrencyRepository(db database.DatabaseInterface, logger log.LoggerInterface) repository.CurrencyRepositoryInterface {
	return &CurrencyRepository{db: db, logger: logger}
}

type CurrencyRepository struct {
	db     database.DatabaseInterface
	logger log.LoggerInterface
}

func (cr *CurrencyRepository) Save(ctx context.Context, currency *model.Currency) (*model.Currency, error) {
	mappedCurrency := mapping.Currency(*currency)
	res := cr.db.GetConnection().Save(&mappedCurrency)
	if res.Error != nil {
		cr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return currency, nil
}

func (cr *CurrencyRepository) GetCurrencyByFigi(ctx context.Context, figi string) (*model.Currency, error) {
	var currency model.Currency

	res := cr.db.GetConnection().Model(&mapping.Currency{}).Where("figi = ?", figi).Find(&currency)
	if res.Error != nil {
		cr.logger.Error(log.LogCategoryDatabase, res.Error.Error(), make(map[string]interface{}))
		return nil, res.Error
	}

	return &currency, nil
}
