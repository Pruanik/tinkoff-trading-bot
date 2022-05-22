package repository

import (
	"context"
	"errors"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
)

func NewCurrencyRepository(db database.DatabaseInterface) repository.CurrencyRepositoryInterface {
	return &CurrencyRepository{db: db}
}

type CurrencyRepository struct {
	db database.DatabaseInterface
}

func (cr *CurrencyRepository) Save(ctx context.Context, currency *model.Currency) (*model.Currency, error) {
	res := cr.db.GetConnection().Save(currency)
	if res.Error != nil {
		return nil, res.Error
	}

	return currency, nil
}

func (cr *CurrencyRepository) GetCurrencyByFigi(ctx context.Context, figi string) (*model.Currency, error) {
	var currency model.Currency

	result := cr.db.GetConnection().Model(&model.Currency{}).Where("figi = ?", figi).Find(&currency)
	if result.Error != nil {
		return nil, errors.New("GetCurrencyByFigi: " + result.Error.Error())
	}

	return &currency, nil
}
