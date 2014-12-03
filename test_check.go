package gose4

import "time"

// Simple check to use for tests around gose4 code
type TestCheck struct{
	TestCheckResult HealthCheckResult
	TestCheckConfig Configuration
}

func NewTestCheck() TestCheck {
	return TestCheck{DefaultTestCheckResult, DefaultTestCheckConfig}
}

var DefaultTestCheckResult = HealthCheckResult{
	Duration:  time.Second * 5,
	Result:    CheckPassed,
	Message:   "It works!",
	LastCheck: time.Unix(1415030870, 0),
}
var DefaultTestCheckConfig = Configuration{
	Severity:     SeverityWarn,
	InitialDelay: time.Second * 0,
	RunDelay:     time.Second,
	Description:  "Example config",
}

func (t TestCheck) Run() HealthCheckResult {
	return t.TestCheckResult
}
func (t TestCheck) Configuration() Configuration {
	return t.TestCheckConfig
}

