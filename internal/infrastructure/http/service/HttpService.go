package service

import (
	"github.com/gin-gonic/gin"
)

func NewHttpService() HttpServiceInterface {
	return &HttpService{}
}

type HttpService struct{}

func (hs HttpService) GetQueryParams(ctx *gin.Context, paramKey string) *string {
	paramValue, isParamExist := ctx.GetQuery(paramKey)

	var result *string
	if isParamExist {
		result = &paramValue
	}

	return result
}
