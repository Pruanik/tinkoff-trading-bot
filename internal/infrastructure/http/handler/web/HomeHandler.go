package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

type HomeHandler struct{}

func (hh *HomeHandler) Handle(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		"",
	)
}
