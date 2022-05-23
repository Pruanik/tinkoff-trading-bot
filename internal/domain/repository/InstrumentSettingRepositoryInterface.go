package repository

import (
	"context"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type InstrumentSettingRepositoryInterface interface {
	Save(ctx context.Context, instrumentSetting *model.InstrumentSetting) (*model.InstrumentSetting, error)

	GetInstrumentsSettings(ctx context.Context) ([]model.InstrumentSettingWithName, error)

	GetInstrumentSettingByFigi(ctx context.Context, figi string) (*model.InstrumentSetting, error)

	Update(ctx context.Context, instrumentSetting *model.InstrumentSetting) (*model.InstrumentSetting, error)
}
