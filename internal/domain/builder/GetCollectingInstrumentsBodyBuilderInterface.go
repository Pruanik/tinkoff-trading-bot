package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/httpresponse/api/item"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type GetCollectingInstrumentsBodyBuilderInterface interface {
	CreateBody(instrumentsSettings []model.InstrumentSettingWithName) []item.GetCollectingInstrumentsResponseBody
}
