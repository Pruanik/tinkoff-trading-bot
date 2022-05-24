package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/httpresponse/api/item"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

func NewGetCandlesChartBodyBuilder() builder.GetCandlesChartBodyBuilderInterface {
	return &GetCandlesChartBodyBuilder{}
}

type GetCandlesChartBodyBuilder struct{}

func (gccbb GetCandlesChartBodyBuilder) CreateBody(candles []model.Candle) []item.GetCandlesChartResponseBody {
	var body []item.GetCandlesChartResponseBody

	for i := 0; i < len(candles); i++ {
		item := item.GetCandlesChartResponseBody{
			High:      candles[i].High,
			Timestamp: candles[i].Timestamp.Unix(),
		}
		body = append(body, item)
	}

	return body
}
