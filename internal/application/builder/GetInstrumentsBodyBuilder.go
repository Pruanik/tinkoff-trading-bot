package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/httpresponse/api/item"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

func NewGetInstrumentsBodyBuilder() builder.GetInstrumentsBodyBuilderInterface {
	return &GetInstrumentsBodyBuilder{}
}

type GetInstrumentsBodyBuilder struct{}

func (gibb GetInstrumentsBodyBuilder) CreateBody(instruments []model.Instrument) []item.GetInstrumentsResponseBody {
	var body []item.GetInstrumentsResponseBody

	for i := 0; i < len(instruments); i++ {
		item := item.GetInstrumentsResponseBody{
			Id:        instruments[i].Id,
			Figi:      instruments[i].Figi,
			Name:      instruments[i].Name,
			Type:      instruments[i].Type,
			CreatedAt: instruments[i].CreatedAt,
		}
		body = append(body, item)
	}

	return body
}
