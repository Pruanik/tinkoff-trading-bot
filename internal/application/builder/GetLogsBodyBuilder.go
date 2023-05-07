package builder

import (
	"encoding/json"

	"github.com/Pruanik/tinkoff-trading-bot/internal/application/http/api/response"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/model"
)

func NewGetLogsBodyBuilder() builder.GetLogsBodyBuilderInterface {
	return &GetLogsBodyBuilder{}
}

type GetLogsBodyBuilder struct{}

func (glbb GetLogsBodyBuilder) CreateBody(logs []model.Log) []response.GetLogsResponseBody {
	var body []response.GetLogsResponseBody

	for i := 0; i < len(logs); i++ {
		var context map[string]string
		json.Unmarshal(logs[i].Context, &context)

		item := response.GetLogsResponseBody{
			Id:        logs[i].Id,
			Category:  logs[i].Category,
			Level:     logs[i].Level,
			Message:   logs[i].Message,
			Context:   context,
			CreatedAt: logs[i].CreatedAt,
		}
		body = append(body, item)
	}

	return body
}
