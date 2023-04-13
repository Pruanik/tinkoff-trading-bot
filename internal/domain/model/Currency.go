package model

import "time"

type Currency struct {
	Id                     int64
	Figi                   string
	Ticker                 string
	ClassCode              string
	Isin                   string
	Lot                    int32
	Currency               string
	Name                   string
	Exchange               string
	OtcFlag                bool
	BuyAvailableFlag       bool
	SellAvailableFlag      bool
	IsoCurrencyName        string
	MinPriceIncrementUnits int64
	MinPriceIncrementNano  int32
	ApiTradeAvailableFlag  bool
	CreatedAt              time.Time
	UpdatedAt              time.Time
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
