package time

import "time"

func SecondsDuration(s int) time.Duration {
	return time.Duration(time.Second * time.Duration(s))
}
