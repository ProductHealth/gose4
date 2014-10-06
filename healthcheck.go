package gose4

import (
	"fmt"
	"github.com/golang/glog"
	"time"
)

type Severity int
type Result int

const (
	SeverityNotice Severity = iota + 1
	SeverityWarn
	SeverityFatal
)

const (
	CheckPassed Result = iota + 1
	CheckFailed
	CheckRunning
)

type Configuration struct {
	Severity     Severity
	InitialDelay time.Duration
	RunDelay     time.Duration
	Description  string
}

//initialDelay and period are integer representin seconds
func NewConfiguration(description string, severity Severity, initialDelay, period int) Configuration {
	return Configuration{
		Severity:     SeverityWarn,
		InitialDelay: time.Second * time.Duration(initialDelay),
		RunDelay:     time.Second * time.Duration(period),
		Description:  description}
}

type HealthCheckResult struct {
	Duration  time.Duration
	Result    Result
	Message   string
	LastCheck time.Time
}

func (r HealthCheckResult) DurationMillis() int64 {
	return r.Duration.Nanoseconds() / 1000000
}

func (r HealthCheckResult) String() string {
	return fmt.Sprintf("HealthCheckResult : [duration=%vms], [result=%v] : %v", r.DurationMillis(), r.Result, r.Message)
}

func (r Result) String() string {
	switch r {
	case CheckPassed:
		return "passed"
	case CheckFailed:
		return "failed"
	case CheckRunning:
		return "running"
	default:
		return "unknown"
	}
}

type HealthCheck interface {
	Run() HealthCheckResult
	Configuration() Configuration
}

// Default return helper methods
func Ok(d time.Duration) HealthCheckResult {
	return Passed(d, "OK")
}
func Passed(d time.Duration, m string) HealthCheckResult {
	return HealthCheckResult{d, CheckPassed, m, time.Now().UTC()}
}

func Failed(d time.Duration, m string) HealthCheckResult {
	return HealthCheckResult{d, CheckFailed, m, time.Now().UTC()}
}


type HealthcheckService struct {
	healthchecks []HealthCheck
	results      map[HealthCheck]HealthCheckResult
}

func CreateHealthcheckService() *HealthcheckService {
	return &HealthcheckService{[]HealthCheck{}, make(map[HealthCheck]HealthCheckResult)}
}

func (se4server HealthcheckService) GetResults() map[HealthCheck]HealthCheckResult {
	return se4server.results
}

func (se4service HealthcheckService) RegisterHealthcheck(h HealthCheck) {
	glog.V(0).Infof("Registering healthcheck '%v'", h.Configuration().Description)
	se4service.healthchecks = append(se4service.healthchecks, h)
	go se4service.runHealthcheck(h)
}

func (se4server HealthcheckService) runHealthcheck(h HealthCheck) {
	// Wait for initial delay to pass
	glog.V(1).Infof("Waiting %v before executing %v", h.Configuration().InitialDelay, h.Configuration().Description)
	time.Sleep(h.Configuration().InitialDelay)
	for {
		glog.V(0).Infof("Executing healthcheck %v", h.Configuration().Description)
		result := h.Run()
		glog.V(1).Infof("Healthcheck returned result : %v", result)
		se4server.results[h] = result
		time.Sleep(h.Configuration().RunDelay)
	}
}
