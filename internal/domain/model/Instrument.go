package model

import "time"

type Instrument struct {
	Id        int64
	Figi      string
	Name      string
	SectorId  int64
	Type      string
	CreatedAt time.Time
}

func NewInstrument(figi string, name string, sectorId int64, instrumentType string) *Instrument {
	return &Instrument{Figi: figi, Name: name, SectorId: sectorId, Type: instrumentType}
}
