package model

type InstrumentSector struct {
	Id   int64
	Code string
	Name string
}

func NewInstrumentSector(code string, name string) *InstrumentSector {
	return &InstrumentSector{Code: code, Name: name}
}
