package api

import "github.com/gin-gonic/gin"

func NewApiRouter() *ApiRouter {
	router := ApiRouter{}
	return &router
}

type ApiRouter struct {
	connection string
	//handler *api.HomeHandler
}

func (ar *ApiRouter) AddGroup(apiRouterGroup *gin.RouterGroup) {
	//apiRouterGroup.GET("/", ar.handler.handle)
}
