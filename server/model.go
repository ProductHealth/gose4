package server

import (
	"time"
)

type Status struct {
	ArtifactId         string  `json:"artifact_id"`
	BuildNumber        string  `json:"build_number"`
	BuildMachine       string  `json:"build_machine"`
	BuildBy            string  `json:"build_by"`
	BuildWhen          string  `json:"build_when"` //ISO 8601 Representation
	CompilerVersion    string  `json:"compiler_version"`
	CurrentTime        string  `json:"current_time"` //ISO 8601 Representation
	GitSha             string  `json:"git_sha"`
	MachineName        string  `json:"machine_name"`
	OsArch             string  `json:"os_arch"`
	OsName             string  `json:"os_name"`
	OsVersion          string  `json:"os_version"`
	RunbookUri         string  `json:"runbook_uri"`
	UpDuration         string  `json:"up_duration"`
	UpSince            string  `json:"up_since"` //ISO 8601 Representation
	Version            string  `json:"version"`
	OsLoad             *string `json:"os_avgload,omitempty"`
	OsNumberProcessors *int    `json:"os_numprocessors,omitempty"`
}

func (s *Status) SetBuildWhen(t *time.Time) {
	s.BuildWhen = timeToIso8601(t.UTC())
}
func (s *Status) SetCurrentTime(t *time.Time) {
	s.CurrentTime = timeToIso8601(t.UTC())
}

type HealthCheck struct {
	ReportAsOf     string            `json:"report_as_of"`    //ISO 8601 Representation
	ReportDuration string            `json:"report_duration"` //
	Tests          []HealthCheckTest `json:"tests"`
}
type HealthCheckTest struct {
	DurationMillis int64  `json:"duration_millis"`
	TestName       string `json:"test_name"`
	TestResult     string `json:"rest_result"`
	TestedAt       string `json:"tested_at"` //ISO 8601 Representation
}
