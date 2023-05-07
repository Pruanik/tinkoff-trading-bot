package response

import "time"

type GetLogsResponseBody struct {
	Id        int64
	Category  string
	Level     string
	Message   string
	Context   interface{}
	CreatedAt time.Time
}
