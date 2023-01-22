package model

import "time"

type Log struct {
	Id        int64
	Category  string
	Level     string
	Message   string
	Context   []byte
	CreatedAt time.Time
}

func NewLogModel(category string, level string, context []byte) *Log {
	return &Log{Category: category, Level: level, Context: context}
}

func (l *Log) TableName() string {
	return "logs"
}
