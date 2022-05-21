package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/httpresponse/api/response"
)

type HttpResponseBuilderInterface interface {
	BuildErrorResponse(message string) response.Error

	BuildSuccessResponse(body response.ResponseBody) response.Success
}
