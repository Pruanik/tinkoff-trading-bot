package api

import "github.com/gin-gonic/gin"

func NewApiRouter() *ApiRouter {
	router := ApiRouter{}
	return &router
}

type ApiRouter struct {
	//handler *api.HomeHandler
}

func (ar *ApiRouter) Create(router *gin.Engine) {
	//apiRouterGroup.GET("/", ar.handler.handle)
}
