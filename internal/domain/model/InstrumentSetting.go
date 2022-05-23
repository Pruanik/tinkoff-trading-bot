package model

import "time"

type InstrumentSetting struct {
	Id               int64     `gorm:"primarykey;column:id"`
	Figi             string    `gorm:"column:figi;type:varchar(255);not null"`
	IsDataCollecting bool      `gorm:"column:is_data_collecting;not null"`
	CreatedAt        time.Time `gorm:"type:timestamp without time zone;not null"`
	UpdatedAt        time.Time `gorm:"type:timestamp without time zone;not null"`
}

func NewInstrumentSetting(figi string, isDataCollecting bool) *InstrumentSetting {
	return &InstrumentSetting{Figi: figi, IsDataCollecting: isDataCollecting}
}

func (is *InstrumentSetting) TableName() string {
	return "instrument_settings"
}
