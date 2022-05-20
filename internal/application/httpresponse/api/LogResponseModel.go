package api

import (
	"time"
)

type LogResponseModel struct {
	RequestId string
	Time      time.Time
	//Body      []model.LogModel
}
