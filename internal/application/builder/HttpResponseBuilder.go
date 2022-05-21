package builder

import (
	"time"

	"github.com/Pruanik/tinkoff-trading-bot/internal/application/httpresponse/api/response"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
)

func NewHttpResponseBuilder() builder.HttpResponseBuilderInterface {
	return &HttpResponseBuilder{}
}

type HttpResponseBuilder struct{}

func (hrb HttpResponseBuilder) BuildErrorResponse(message string) response.Error {
	responseStatus := response.Status{Status: response.StatusError, Message: message}
	errorResponse := response.Error{Status: responseStatus, Time: time.Now()}

	return errorResponse
}

func (hrb HttpResponseBuilder) BuildSuccessResponse(body response.ResponseBody) response.Success {
	responseStatus := response.Status{Status: response.StatusSuccess, Message: ""}
	successResponse := response.Success{Status: responseStatus, Body: body, Time: time.Now()}

	return successResponse
}
