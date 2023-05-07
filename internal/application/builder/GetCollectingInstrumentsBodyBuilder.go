package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/http/api/response"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

func NewGetCollectingInstrumentsBodyBuilder() builder.GetCollectingInstrumentsBodyBuilderInterface {
	return &GetCollectingInstrumentsBodyBuilder{}
}

type GetCollectingInstrumentsBodyBuilder struct{}

func (gcilbb GetCollectingInstrumentsBodyBuilder) CreateBody(instrumentsSettings []model.InstrumentSettingWithName) []response.GetCollectingInstrumentsResponseBody {
	var body []response.GetCollectingInstrumentsResponseBody

	for i := 0; i < len(instrumentsSettings); i++ {
		item := response.GetCollectingInstrumentsResponseBody{
			Figi:             instrumentsSettings[i].Figi,
			Name:             instrumentsSettings[i].Name,
			IsDataCollecting: instrumentsSettings[i].IsDataCollecting,
			CreatedAt:        instrumentsSettings[i].CreatedAt,
		}
		body = append(body, item)
	}

	return body
}
