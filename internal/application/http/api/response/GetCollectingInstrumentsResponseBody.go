package response

import "time"

type GetCollectingInstrumentsResponseBody struct {
	Figi             string
	Name             string
	IsDataCollecting bool
	CreatedAt        time.Time
}
