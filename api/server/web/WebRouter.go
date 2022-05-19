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

func (wr *WebRouter) Create(router *gin.Engine) {
	router.LoadHTMLGlob("./web/dist/*.html")
	router.Static("/assets", "./web/dist/assets")
	router.Static("/static", "./web/static")

	webRouterGroup := router.Group("/")
	webRouterGroup.GET("/", wr.handler.Handle)
}
