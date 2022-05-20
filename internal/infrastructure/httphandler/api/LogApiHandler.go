package api

import (
	"net/http"
	"time"

	"github.com/Pruanik/tinkoff-trading-bot/internal/application/httpresponse/api"
	"github.com/gin-gonic/gin"
)

func NewLogApiHandler() *LogApiHandler {
	handler := LogApiHandler{}
	return &handler
}

type LogApiHandler struct{}

func (lah LogApiHandler) Handle(ctx *gin.Context) {
	logResponse := api.LogResponseModel{RequestId: "123", Time: time.Now()}
	ctx.JSONP(http.StatusOK, logResponse)
}
