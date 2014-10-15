package heartbeat

import (
	"fmt"
	"github.com/ProductHealth/gose4"
	"time"
)

const (
	nilDuration = time.Second * 0
)

//Heartbeat check, has to triggered within a duration in order to pass
type Check struct {
	maxDuration     time.Duration
	lastTrigger     *time.Time
	config          gose4.Configuration
}

func New(maxDuration time.Duration, config gose4.Configuration) *Check {
	return &Check{maxDuration, nil, config}
}

func (hc *Check) Run() gose4.HealthCheckResult {
	if hc.lastTrigger != nil && time.Since(*hc.lastTrigger) > hc.maxDuration {
		return gose4.Failed(nilDuration, fmt.Sprintf("Last trigger took plate %v, longer than allowed %v ago", *hc.lastTrigger, hc.maxDuration))
	} else {
		return gose4.Ok(nilDuration)
	}
}

func (hc *Check) Configuration() gose4.Configuration {
	return hc.config
}

func (hc *Check) Trigger() {
	now := time.Now()
	hc.lastTrigger = &now
}
