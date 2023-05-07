package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/http/api/response"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type GetInstrumentsBodyBuilderInterface interface {
	CreateBody(instruments []model.Instrument) []response.GetInstrumentsResponseBody
}
