package mapping

import "time"

type Log struct {
	Id        int64     `gorm:"primarykey;column:id"`
	Category  string    `gorm:"column:category;type:varchar(255);not null"`
	Level     string    `gorm:"column:level;type:varchar(255);not null"`
	Message   string    `gorm:"column:message;type:varchar(255);not null"`
	Context   []byte    `gorm:"type:jsonb;column:context;not null"`
	CreatedAt time.Time `gorm:"type:timestamp without time zone;not null"`
}

func (l *Log) TableName() string {
	return "logs"
}
