package web

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/httphandler/web"
	"github.com/gin-gonic/gin"
)

func NewWebRouter(handler *web.HomeHandler) *WebRouter {
	router := WebRouter{handler: handler}
	return &router
}

type WebRouter struct {
	handler *web.HomeHandler
}

func (wr *WebRouter) AddGroup(webRouterGroup *gin.RouterGroup) {
	webRouterGroup.GET("/", wr.handler.Handle)
}
