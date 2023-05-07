package response

import "time"

type GetInstrumentsResponseBody struct {
	Id        int64
	Figi      string
	Name      string
	Type      string
	CreatedAt time.Time
}
