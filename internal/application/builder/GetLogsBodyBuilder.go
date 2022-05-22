package builder

import (
	"encoding/json"

	"github.com/Pruanik/tinkoff-trading-bot/internal/application/httpresponse/api/item"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

func NewGetLogsBodyBuilder() builder.GetLogsBodyBuilderInterface {
	return &GetLogsBodyBuilder{}
}

type GetLogsBodyBuilder struct{}

func (glbb GetLogsBodyBuilder) CreateBody(logs []model.Log) []item.GetLogsResponseBody {
	var body []item.GetLogsResponseBody

	for i := 0; i < len(logs); i++ {
		var context map[string]string
		json.Unmarshal(logs[i].Context, &context)

		item := item.GetLogsResponseBody{
			Id:        logs[i].Id,
			Category:  logs[i].Category,
			Level:     logs[i].Level,
			Context:   context,
			CreatedAt: logs[i].CreatedAt,
		}
		body = append(body, item)
	}

	return body
}
