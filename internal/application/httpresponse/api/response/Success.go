package response

import (
	"time"
)

type ResponseBody interface{}

type Success struct {
	Status Status
	Body   ResponseBody
	Time   time.Time
}
