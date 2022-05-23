package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Candle struct {
	Id        int64           `gorm:"primarykey;column:id"`
	Figi      string          `gorm:"column:figi;type:varchar(255);not null"`
	Open      decimal.Decimal `gorm:"column:open;type:decimal;not null"`
	High      decimal.Decimal `gorm:"column:high;type:decimal;not null"`
	Low       decimal.Decimal `gorm:"column:low;type:decimal;not null"`
	Close     decimal.Decimal `gorm:"column:close;type:decimal;not null"`
	Volume    int64           `gorm:"column:volume;type:bigint;not null"`
	Timestamp time.Time       `gorm:"type:timestamp without time zone;not null"`
	CreatedAt time.Time       `gorm:"type:timestamp without time zone;not null"`
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
