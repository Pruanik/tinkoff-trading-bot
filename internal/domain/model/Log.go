package model

import "time"

type Log struct {
	Id        int64     `gorm:"primarykey;column:id"`
	Category  string    `gorm:"column:category;type:varchar(255);not null"`
	Level     string    `gorm:"column:level;type:varchar(255);not null"`
	Context   []byte    `gorm:"type:jsonb;column:context;not null"`
	CreatedAt time.Time `gorm:"type:timestamp without time zone;not null"`
}

func NewLogModel(category string, level string, context []byte) *Log {
	return &Log{Category: category, Level: level, Context: context}
}

func (l *Log) TableName() string {
	return "logs"
}
