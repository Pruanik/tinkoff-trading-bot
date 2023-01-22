package mapping

import "time"

type Task struct {
	Id            int64     `gorm:"primarykey;column:id"`
	Label         string    `gorm:"column:label;type:varchar(255);not null"`
	Arguments     []byte    `gorm:"type:jsonb;column:arguments;not null"`
	Status        int32     `gorm:"column:status;not null"`
	IsRescheduled bool      `gorm:"column:is_rescheduled;not null"`
	TimeShift     *int64    `gorm:"column:time_shift"`
	ExecutedAt    time.Time `gorm:"column:is_rescheduled;type:timestamp without time zone;not null"`
	CreatedAt     time.Time `gorm:"type:timestamp without time zone;not null"`
	UpdatedAt     time.Time `gorm:"type:timestamp without time zone;not null"`
}

func (t *Task) TableName() string {
	return "tasks"
}
