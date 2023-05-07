package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/http/api/response"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type GetInstrumentSectorsBodyBuilderInterface interface {
	CreateBody(instrumentSector []model.InstrumentSector) []response.GetInstrumentSectorsResponseBody
}
