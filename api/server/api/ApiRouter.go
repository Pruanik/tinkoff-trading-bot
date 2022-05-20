package api

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/httphandler/api"
	"github.com/gin-gonic/gin"
)

func NewApiRouter(logApiHandler *api.LogApiHandler) *ApiRouter {
	router := ApiRouter{logApiHandler: logApiHandler}
	return &router
}

type ApiRouter struct {
	logApiHandler *api.LogApiHandler
}

func (ar *ApiRouter) Create(router *gin.Engine) {
	apiRouterGroup := router.Group("/api/")
	apiRouterGroup.GET("/getLogs", ar.logApiHandler.Handle)
}
