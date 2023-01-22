package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Candle struct {
	Id        int64
	Figi      string
	Open      decimal.Decimal
	High      decimal.Decimal
	Low       decimal.Decimal
	Close     decimal.Decimal
	Volume    int64
	Timestamp time.Time
	CreatedAt time.Time
}

func NewCandle(
	figi string,
	open decimal.Decimal,
	high decimal.Decimal,
	low decimal.Decimal,
	close decimal.Decimal,
	volume int64,
	timestamp time.Time,
) *Candle {
	return &Candle{
		Figi:      figi,
		Open:      open,
		High:      high,
		Low:       low,
		Close:     close,
		Volume:    volume,
		Timestamp: timestamp,
	}
}

func (c *Candle) TableName() string {
	return "candles"
}
