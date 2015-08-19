package time

import "time"

type Clock struct {
	fixedTime *time.Time
}

func (clock *Clock) Now() time.Time {
	switch {
	case clock.fixedTime != nil: return *clock.fixedTime
	default: return time.Now()
	}
}

func NewClock() Clock {
	return Clock{nil}
}

func (clock *Clock) SetTime(t *time.Time) {
	clock.fixedTime = t
}

func (clock *Clock) Lock() {
	t := time.Now()
	clock.SetTime(&t)
}
