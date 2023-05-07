package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/http/api/response"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

func NewGetInstrumentSectorsBodyBuilder() builder.GetInstrumentSectorsBodyBuilderInterface {
	return &GetInstrumentSectorsBodyBuilder{}
}

type GetInstrumentSectorsBodyBuilder struct{}

func (gisbb GetInstrumentSectorsBodyBuilder) CreateBody(instrumentSectors []model.InstrumentSector) []response.GetInstrumentSectorsResponseBody {
	var body []response.GetInstrumentSectorsResponseBody

	for i := 0; i < len(instrumentSectors); i++ {
		item := response.GetInstrumentSectorsResponseBody{
			Id:   instrumentSectors[i].Id,
			Code: instrumentSectors[i].Code,
			Name: instrumentSectors[i].Name,
		}
		body = append(body, item)
	}

	return body
}
