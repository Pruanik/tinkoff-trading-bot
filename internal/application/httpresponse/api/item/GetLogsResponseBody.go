package item

import "time"

type GetLogsResponseBody struct {
	Id        int64
	Category  string
	Level     string
	Context   interface{}
	CreatedAt time.Time
}
