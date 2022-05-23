package dateoperation

import "time"

func NewDateOperation() DateOperationInterface {
	return &DateOperation{}
}

type DateOperation struct{}

func (do DateOperation) MonthsCountSince(timeInPast time.Time) int {
	now := time.Now()
	months := 0
	month := timeInPast.Month()
	for timeInPast.Before(now) {
		timeInPast = timeInPast.Add(time.Hour * 24)
		nextMonth := timeInPast.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}

type DateOperationInterface interface {
	MonthsCountSince(timeInPast time.Time) int
}
