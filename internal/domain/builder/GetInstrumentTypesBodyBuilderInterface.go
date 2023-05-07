package builder

import "github.com/Pruanik/tinkoff-trading-bot/internal/application/http/api/response"

type GetInstrumentTypesBodyBuilderInterface interface {
	CreateBody(instrumentTypes []string) []response.GetInstrumentTypesResponseBody
}
