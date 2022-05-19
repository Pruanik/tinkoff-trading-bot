package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHomeHandler() *HomeHandler {
	handler := HomeHandler{}
	return &handler
}

type HomeHandler struct{}

func (hh *HomeHandler) Handle(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		"",
	)
}
