package api

import (
	"net/http"

	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/Pruanik/tinkoff-trading-bot/internal/domain/builder"
	"github.com/gin-gonic/gin"
)

func NewSystemApiHandler(
	httpResponseBuilder builder.HttpResponseBuilderInterface,
	config *configs.Config,
) *SystemApiHandler {
	return &SystemApiHandler{
		httpResponseBuilder: httpResponseBuilder,
		config:              config,
	}
}

type SystemApiHandler struct {
	httpResponseBuilder builder.HttpResponseBuilderInterface
	config              *configs.Config
}

func (sah SystemApiHandler) HandleGetMod(ctx *gin.Context) {
	type getModResult struct {
		Mod string
	}

	responseBody := getModResult{Mod: sah.config.TinkoffInvestConfig.Mod}
	ctx.JSONP(http.StatusOK, sah.httpResponseBuilder.BuildSuccessResponse(responseBody))
}
