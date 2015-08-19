package time

import "time"

type stopWatch struct {
	Start time.Time
}

func NewStopWatch() *stopWatch {
	sw := stopWatch{time.Now()}
	return &sw
}

func (sw stopWatch) GetDuration() time.Duration {
	return time.Now().Sub(sw.Start)
}
func (sw stopWatch) GetMillis() int64 {
	return sw.GetDuration().Nanoseconds() * 1000
}

func (sw stopWatch) WaitUntil(d time.Duration) {
	intervalRemaining := time.Now().Add(d).Sub(time.Now().Add(sw.GetDuration()))
	time.Sleep(intervalRemaining)
}
