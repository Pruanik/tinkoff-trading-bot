package web

import "github.com/gin-gonic/gin"

func NewHomeHandler() *HomeHandler {
	handler := HomeHandler{}
	return &handler
}

type HomeHandler struct{}

func (hh *HomeHandler) Handle(ctx *gin.Context) {
	//handler.ServeHTTP(ctx.Writer, ctx.Request)
}
