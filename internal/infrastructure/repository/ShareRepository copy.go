package repository

import (
	"context"
	"errors"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/repository"
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/database"
)

func NewShareRepository(db database.DatabaseInterface) repository.ShareRepositoryInterface {
	return &ShareRepository{db: db}
}

type ShareRepository struct {
	db database.DatabaseInterface
}

func (sr *ShareRepository) Save(ctx context.Context, share *model.Share) (*model.Share, error) {
	res := sr.db.GetConnection().Save(share)
	if res.Error != nil {
		return nil, res.Error
	}

	return share, nil
}

func (sr *ShareRepository) GetShareByFigi(ctx context.Context, figi string) (*model.Share, error) {
	var share model.Share

	result := sr.db.GetConnection().Model(&model.Share{}).Where("figi = ?", figi).Find(&share)
	if result.Error != nil {
		return nil, errors.New("GetShareByFigi: " + result.Error.Error())
	}

	return &share, nil
}
