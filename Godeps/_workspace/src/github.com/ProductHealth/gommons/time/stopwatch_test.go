package time

import (
	"testing"
	"time"
)

func TestStopWatch(t *testing.T) {
	sw := NewStopWatch()
	time.Sleep(time.Second * 1)
	if sw.GetDuration().Seconds() < 1 {
		t.Error("Stopwatch did not record required duration")
		t.Fail()
	}
	if sw.GetMillis() < 5000 {
		t.Error("Stopwatch did not record required millis")
		t.Fail()
	}

}
