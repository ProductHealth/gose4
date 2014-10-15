package heartbeat

import (
	"github.com/ProductHealth/gose4"
	"github.com/stretchr/testify/assert"
	"time"
	"testing"
)

var (
	oneminute  = time.Minute * 1
	twominutes = time.Minute * 2
)

var config = gose4.NewConfiguration("test", gose4.SeverityWarn, 1, 1)

func TestFailsIfDurationPassed(t *testing.T) {
	twoMinuteAgo := time.Now().Add(-twominutes)
	c := Check{oneminute, &twoMinuteAgo, config}
	res := c.Run()
	assert.Equal(t, gose4.CheckFailed, res.Result)
}

func TestPassesIfDurationWithinWindow(t *testing.T) {
	oneMinuteAgo := time.Now().Add(-oneminute)
	c := Check{twominutes, &oneMinuteAgo, config}
	res := c.Run()
	assert.Equal(t, gose4.CheckPassed, res.Result)
}

func TestFailsAndPassesAfterTrigger(t *testing.T) {
	twoMinuteAgo := time.Now().Add(-twominutes)
	c := Check{oneminute, &twoMinuteAgo, config}
	res := c.Run()
	assert.Equal(t, gose4.CheckFailed, res.Result)
	c.Trigger()
	res = c.Run()
	assert.Equal(t, gose4.CheckPassed, res.Result)

}
