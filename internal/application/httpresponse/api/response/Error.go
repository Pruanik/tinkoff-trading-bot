package response

import (
	"time"
)

type Error struct {
	Status Status
	Time   time.Time
}
