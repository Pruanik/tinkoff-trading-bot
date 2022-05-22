package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/httpresponse/api/item"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

type GetLogsBodyBuilderInterface interface {
	CreateBody(logs []model.Log) []item.GetLogsResponseBody
}
