package model

import "time"

type Currency struct {
	Id                     int64     `gorm:"primarykey;column:id"`
	Figi                   string    `gorm:"column:figi;type:varchar(255);not null"`
	Ticker                 string    `gorm:"column:ticker;type:varchar(255);not null"`
	ClassCode              string    `gorm:"column:class_code;type:varchar(255);not null"`
	Isin                   string    `gorm:"column:isin;type:varchar(255);not null"`
	Lot                    int32     `gorm:"column:lot;not null"`
	Currency               string    `gorm:"column:currency;type:varchar(255);not null"`
	Name                   string    `gorm:"column:name;type:varchar(255);not null"`
	Exchange               string    `gorm:"column:exchange;type:varchar(255);not null"`
	OtcFlag                bool      `gorm:"column:otc_flag;not null"`
	BuyAvailableFlag       bool      `gorm:"column:buy_available_flag;not null"`
	SellAvailableFlag      bool      `gorm:"column:sell_available_flag;not null"`
	IsoCurrencyName        string    `gorm:"column:iso_currency_name;type:varchar(255);not null"`
	MinPriceIncrementUnits int64     `gorm:"column:min_price_increment_units;not null"`
	MinPriceIncrementNano  int32     `gorm:"column:min_price_increment_nano;not null"`
	ApiTradeAvailableFlag  bool      `gorm:"column:api_trade_available_flag;not null"`
	CreatedAt              time.Time `gorm:"type:timestamp without time zone;not null"`
	UpdatedAt              time.Time `gorm:"type:timestamp without time zone;not null"`
}

func NewCurrency(
	figi string,
	ticker string,
	classCode string,
	isin string,
	lot int32,
	currency string,
	name string,
	exchange string,
	otcFlag bool,
	buyAvailableFlag bool,
	sellAvailableFlag bool,
	isoCurrencyName string,
	minPriceIncrementUnits int64,
	minPriceIncrementNano int32,
	apiTradeAvailableFlag bool,
) *Currency {
	return &Currency{
		Figi:                   figi,
		Ticker:                 ticker,
		ClassCode:              classCode,
		Isin:                   isin,
		Lot:                    lot,
		Currency:               currency,
		Name:                   name,
		Exchange:               exchange,
		OtcFlag:                otcFlag,
		BuyAvailableFlag:       buyAvailableFlag,
		SellAvailableFlag:      sellAvailableFlag,
		IsoCurrencyName:        isoCurrencyName,
		MinPriceIncrementUnits: minPriceIncrementUnits,
		MinPriceIncrementNano:  minPriceIncrementNano,
		ApiTradeAvailableFlag:  apiTradeAvailableFlag,
	}
}

func (c *Currency) TableName() string {
	return "currencies"
}
