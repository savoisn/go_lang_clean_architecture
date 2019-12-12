package entity

import (
	"time"
)

type Task struct {
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Type        string
}
