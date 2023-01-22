package model

import "time"

type InstrumentSetting struct {
	Id               int64
	Figi             string
	IsDataCollecting bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewInstrumentSetting(figi string, isDataCollecting bool) *InstrumentSetting {
	return &InstrumentSetting{Figi: figi, IsDataCollecting: isDataCollecting}
}

func (is *InstrumentSetting) TableName() string {
	return "instrument_settings"
}
