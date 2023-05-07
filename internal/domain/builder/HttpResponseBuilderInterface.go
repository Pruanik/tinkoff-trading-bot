package builder

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/application/http/common"
)

type HttpResponseBuilderInterface interface {
	BuildErrorResponse(message string) common.Error

	BuildSuccessResponse(body common.ResponseBody) common.Success
}
