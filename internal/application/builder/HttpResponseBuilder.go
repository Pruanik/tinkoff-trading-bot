package builder

import (
	"time"

	"github.com/Pruanik/tinkoff-trading-bot/internal/application/http/common"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
)

func NewHttpResponseBuilder() builder.HttpResponseBuilderInterface {
	return &HttpResponseBuilder{}
}

type HttpResponseBuilder struct{}

func (hrb HttpResponseBuilder) BuildErrorResponse(message string) common.Error {
	responseStatus := common.Status{Status: common.StatusError, Message: message}
	errorResponse := common.Error{Status: responseStatus, Time: time.Now()}

	return errorResponse
}

func (hrb HttpResponseBuilder) BuildSuccessResponse(body common.ResponseBody) common.Success {
	responseStatus := common.Status{Status: common.StatusSuccess, Message: ""}
	successResponse := common.Success{Status: responseStatus, Body: body, Time: time.Now()}

	return successResponse
}
