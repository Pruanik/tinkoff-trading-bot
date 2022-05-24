package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type GetCandlesChartBodyBuilderInterface interface {
	CreateBody(candles []model.Candle) []interface{}
}
