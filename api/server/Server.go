package server

import (
	"fmt"
	"net/http"

	"github.com/Pruanik/tinkoff-trading-bot/api/server/api"
	"github.com/Pruanik/tinkoff-trading-bot/api/server/web"
	"github.com/Pruanik/tinkoff-trading-bot/configs"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	return router
}

func NewServer(router *gin.Engine, config *configs.Config) *http.Server {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.WebApplicationConfig.Port),
		Handler: router,
	}

	return server
}

func RegisterRoutes(router *gin.Engine, webRouter *web.WebRouter, apiRouter *api.ApiRouter) {
	webRouterGroup := router.Group("/")
	webRouter.AddGroup(webRouterGroup)

	apiRouterGroup := router.Group("/api/")
	apiRouter.AddGroup(apiRouterGroup)
}
