package api

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/httphandler/api"
	"github.com/gin-gonic/gin"
)

func NewApiRouter(
	logApiHandler *api.LogApiHandler,
	instrumentApiHandler *api.InstrumentApiHandler,
	systemApiHandler *api.SystemApiHandler,
	candleApiHandler *api.CandleApiHandler,
) *ApiRouter {
	router := ApiRouter{
		logApiHandler:        logApiHandler,
		instrumentApiHandler: instrumentApiHandler,
		systemApiHandler:     systemApiHandler,
		candleApiHandler:     candleApiHandler,
	}
	return &router
}

type ApiRouter struct {
	logApiHandler        *api.LogApiHandler
	instrumentApiHandler *api.InstrumentApiHandler
	systemApiHandler     *api.SystemApiHandler
	candleApiHandler     *api.CandleApiHandler
}

func (ar *ApiRouter) Create(router *gin.Engine) {
	apiRouterGroup := router.Group("/api/")
	apiRouterGroup.GET("/getLogs", ar.logApiHandler.Handle)
	apiRouterGroup.GET("/getInstruments", ar.instrumentApiHandler.HandleGetInstruments)
	apiRouterGroup.GET("/getCollectingInstruments", ar.instrumentApiHandler.HandleGetCollectingInstruments)
	apiRouterGroup.GET("/setCollectingInstrument", ar.instrumentApiHandler.HandleSetCollectingInstrument)
	apiRouterGroup.GET("/getMod", ar.systemApiHandler.HandleGetMod)
	apiRouterGroup.GET("/getPeriodCandles", ar.candleApiHandler.HandleGetPeriodCandles)
	apiRouterGroup.GET("/getLastCandles", ar.candleApiHandler.HandleGetLastCandles)
}
