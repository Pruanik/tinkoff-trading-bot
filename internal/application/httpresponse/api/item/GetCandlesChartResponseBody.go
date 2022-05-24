package item

import "github.com/shopspring/decimal"

type GetCandlesChartResponseBody struct {
	Timestamp int64
	High      decimal.Decimal
}
