package util

import "time"

type stopWatch struct {
	Start time.Time
}

func CreateStopWatch() *stopWatch {
	sw := stopWatch{time.Now()}
	return &sw
}

func (sw stopWatch) GetDuration() time.Duration {
	return time.Now().Sub(sw.Start)
}
func (sw stopWatch) GetMillis() int64 {
	return sw.GetDuration().Nanoseconds() * 1000
}
