package builder

import "github.com/Pruanik/tinkoff-trading-bot/internal/application/httpresponse/api/item"

type GetInstrumentTypesBodyBuilderInterface interface {
	CreateBody(instrumentTypes []string) []item.GetInstrumentTypesResponseBody
}
