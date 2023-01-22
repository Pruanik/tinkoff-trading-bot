package mapping

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

func (c *Candle) TableName() string {
	return "candles"
}
