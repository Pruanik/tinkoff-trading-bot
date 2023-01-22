package model

import "time"

type Share struct {
	Id                     int64
	Figi                   string
	Ticker                 string
	ClassCode              string
	Isin                   string
	Lot                    int32
	Currency               string
	Name                   string
	Exchange               string
	Sector                 string
	MinPriceIncrementUnits int64
	MinPriceIncrementNano  int32
	ApiTradeAvailableFlag  bool
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func NewShare(
	figi string,
	ticker string,
	classCode string,
	isin string,
	lot int32,
	currency string,
	name string,
	exchange string,
	sector string,
	minPriceIncrementUnits int64,
	minPriceIncrementNano int32,
	apiTradeAvailableFlag bool,
) *Share {
	return &Share{
		Figi:                   figi,
		Ticker:                 ticker,
		ClassCode:              classCode,
		Isin:                   isin,
		Lot:                    lot,
		Currency:               currency,
		Name:                   name,
		Exchange:               exchange,
		Sector:                 sector,
		MinPriceIncrementUnits: minPriceIncrementUnits,
		MinPriceIncrementNano:  minPriceIncrementNano,
		ApiTradeAvailableFlag:  apiTradeAvailableFlag,
	}
}

func (s *Share) TableName() string {
	return "shares"
}
