package healthcheck

import (
	"fmt"
	"time"
)

type Severity int
type Status int

const (
	SeverityNotice Severity = iota + 1
	SeverityWarn
	SeverityFatal
)
const (
	StatusPassed Status = iota + 1
	StatusFailed
	StatusRunning
)

type HealthCheckConfiguration struct {
	Severity     Severity
	InitialDelay time.Duration
	RunDelay     time.Duration
	Description  string
}

//initialDelay and period are integer representin seconds
func NewConfiguration(description string, severity Severity, initialDelay, period int) HealthCheckConfiguration {
	return HealthCheckConfiguration{
		Severity: SeverityWarn,
		InitialDelay: time.Second * time.Duration(initialDelay),
		RunDelay: time.Second * time.Duration(period),
		Description: description}
}

type HealthCheckResult struct {
	Duration  time.Duration
	Status    Status
	Message   string
	LastCheck time.Time
}

func (r HealthCheckResult) DurationMillis() int64 {
	return r.Duration.Nanoseconds() / 1000000
}

func (r HealthCheckResult) String() string {
	return fmt.Sprintf("HealthCheckResult : [duration=%vms], [result=%v] : %v", r.DurationMillis(), r.Status, r.Message)
}

func (r Status) String() string {
	switch r {
	case StatusPassed:
		return "passed"
	case StatusFailed:
		return "failed"
	case StatusRunning:
		return "running"
	default:
		return "unknown"
	}
}

type HealthCheck interface {
	Run() HealthCheckResult
	Configuration() HealthCheckConfiguration
}

// Default return helper methods
func Ok(d time.Duration) HealthCheckResult {
	return Passed(d, "OK")
}
func Passed(d time.Duration, m string) HealthCheckResult {
	return HealthCheckResult{d, StatusPassed, m, time.Now().UTC()}
}

func Failed(d time.Duration, m string) HealthCheckResult {
	return HealthCheckResult{d, StatusFailed, m, time.Now().UTC()}
}
