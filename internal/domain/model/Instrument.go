package model

import "time"

type Instrument struct {
	Id        int64
	Figi      string
	Name      string
	Type      string
	CreatedAt time.Time
}

func NewInstrument(figi string, name string, instrumentType string) *Instrument {
	return &Instrument{Figi: figi, Name: name, Type: instrumentType}
}

func (i *Instrument) TableName() string {
	return "instruments"
}
