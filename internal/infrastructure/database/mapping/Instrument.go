package mapping

import "time"

type Instrument struct {
	Id        int64     `gorm:"primarykey;column:id"`
	Figi      string    `gorm:"column:figi;type:varchar(255);not null"`
	Name      string    `gorm:"column:name;type:varchar(255);not null"`
	Type      string    `gorm:"column:type;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp without time zone;not null"`
}

func (i *Instrument) TableName() string {
	return "instruments"
}
