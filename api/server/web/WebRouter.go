package web

import (
	"github.com/Pruanik/tinkoff-trading-bot/internal/infrastructure/httphandler/web"
	"github.com/gin-gonic/gin"
)

func NewWebRouter(homeWebHandler *web.HomeHandler, pageNotFoundHandler *web.PageNotFoundHandler) *WebRouter {
	router := WebRouter{
		homeWebHandler:      homeWebHandler,
		pageNotFoundHandler: pageNotFoundHandler,
	}
	return &router
}

type WebRouter struct {
	homeWebHandler      *web.HomeHandler
	pageNotFoundHandler *web.PageNotFoundHandler
}

func (wr *WebRouter) Create(router *gin.Engine) {
	router.LoadHTMLGlob("./web/dist/*.html")
	router.Static("/assets", "./web/dist/assets")
	router.Static("/static", "./web/static")

	router.NoRoute(wr.pageNotFoundHandler.Handle)

	webRouterGroup := router.Group("/")
	webRouterGroup.GET("/", wr.homeWebHandler.Handle)
}
