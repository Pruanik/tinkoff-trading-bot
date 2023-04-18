package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/httpresponse/api/item"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NewGetInstrumentTypesBodyBuilder() builder.GetInstrumentTypesBodyBuilderInterface {
	return &GetInstrumentTypesBodyBuilder{}
}

type GetInstrumentTypesBodyBuilder struct{}

func (gitbb GetInstrumentTypesBodyBuilder) CreateBody(instrumentTypes []string) []item.GetInstrumentTypesResponseBody {
	var body []item.GetInstrumentTypesResponseBody

	for i := 0; i < len(instrumentTypes); i++ {
		item := item.GetInstrumentTypesResponseBody{
			Code: instrumentTypes[i],
			Name: cases.Title(language.Und, cases.NoLower).String(instrumentTypes[i]),
		}
		body = append(body, item)
	}

	return body
}
