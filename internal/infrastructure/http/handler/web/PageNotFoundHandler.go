package web

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func NewPageNotFoundHandler() *PageNotFoundHandler {
	return &PageNotFoundHandler{}
}

type PageNotFoundHandler struct{}

func (pnfh *PageNotFoundHandler) Handle(ctx *gin.Context) {
	location := url.URL{Path: "/"}
	ctx.Redirect(http.StatusFound, location.RequestURI())
}
