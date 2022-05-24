package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

func NewGetCandlesChartBodyBuilder() builder.GetCandlesChartBodyBuilderInterface {
	return &GetCandlesChartBodyBuilder{}
}

type GetCandlesChartBodyBuilder struct{}

func (gccbb GetCandlesChartBodyBuilder) CreateBody(candles []model.Candle) []interface{} {
	var body []interface{}

	for i := 0; i < len(candles); i++ {
		var item [2]interface{}
		item[0] = candles[i].Timestamp.Unix()
		item[1] = candles[i].High
		body = append(body, item)
	}

	return body
}
