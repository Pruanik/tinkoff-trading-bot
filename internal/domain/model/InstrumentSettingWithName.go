package model

import "time"

type InstrumentSettingWithName struct {
	Figi             string
	Name             string
	IsDataCollecting bool
	CreatedAt        time.Time
}

func NewInstrumentSettingWithName(figi string, name string, isDataCollecting bool, createdAt time.Time) *InstrumentSettingWithName {
	return &InstrumentSettingWithName{Figi: figi, Name: name, IsDataCollecting: isDataCollecting, CreatedAt: createdAt}
}
