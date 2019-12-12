package entity

import (
	"time"
)

type Planning struct {
	Agent     Agent
	StartDate time.Time
	EndDate   time.Time
	Tasks     []Task
}
