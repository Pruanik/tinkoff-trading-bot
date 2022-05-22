package api

import (
	"net/http"

	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/gin-gonic/gin"
)

func NewLogApiHandler(httpResponseBuilder builder.HttpResponseBuilderInterface) *LogApiHandler {
	return &LogApiHandler{httpResponseBuilder: httpResponseBuilder}
}

type LogApiHandler struct {
	httpResponseBuilder builder.HttpResponseBuilderInterface
}

func (lah LogApiHandler) Handle(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, lah.httpResponseBuilder.BuildErrorResponse("error"))
}
