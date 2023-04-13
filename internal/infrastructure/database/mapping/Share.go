package mapping

import "time"

type Share struct {
	Id                     int64     `gorm:"primarykey;column:id"`
	Figi                   string    `gorm:"column:figi;type:varchar(255);not null;unique"`
	Ticker                 string    `gorm:"column:ticker;type:varchar(255);not null"`
	ClassCode              string    `gorm:"column:class_code;type:varchar(255);not null"`
	Isin                   string    `gorm:"column:isin;type:varchar(255);not null"`
	Lot                    int32     `gorm:"column:lot;not null"`
	Currency               string    `gorm:"column:currency;type:varchar(255);not null"`
	Name                   string    `gorm:"column:name;type:varchar(255);not null"`
	Exchange               string    `gorm:"column:exchange;type:varchar(255);not null"`
	MinPriceIncrementUnits int64     `gorm:"column:min_price_increment_units;not null"`
	MinPriceIncrementNano  int32     `gorm:"column:min_price_increment_nano;not null"`
	ApiTradeAvailableFlag  bool      `gorm:"column:api_trade_available_flag;not null"`
	CreatedAt              time.Time `gorm:"type:timestamp without time zone;not null"`
	UpdatedAt              time.Time `gorm:"type:timestamp without time zone;not null"`
}

func (s *Share) TableName() string {
	return "shares"
}
