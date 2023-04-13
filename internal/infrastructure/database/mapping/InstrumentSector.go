package mapping

type InstrumentSector struct {
	Id   int64  `gorm:"primarykey;column:id"`
	Code string `gorm:"column:code;type:varchar(255);not null"`
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

func (i *InstrumentSector) TableName() string {
	return "instrument_sectors"
}
