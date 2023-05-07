package service

import "github.com/gin-gonic/gin"

type HttpServiceInterface interface {
	GetQueryParams(ctx *gin.Context, paramKey string) *string
}
