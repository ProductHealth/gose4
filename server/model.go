package server

type Status struct {
	ArtifactId      string
	BuildNumber     string
	BuildMachine    string
	BuildBy         string
	BuildWhen       string //ISO 8601 Representation
	CompilerVersion string
	CurrentTime     string //ISO 8601 Representation
	GitSha          string
	MachineName     string
	OsArch          string
	OsName          string
	OsVersion       string
	RunbookUri      string
	UpDuration      string
	UpSince         string //ISO 8601 Representation
	Version         string
	OsLoad          *string
	OsNumberProcessors *int
}

type HealthCheck struct {
	ReportAsOf      string `json:"report_as_of"`//ISO 8601 Representation
	ReportDuration  string `json:"report_duration"`//
	Tests           []HealthCheckTest `json:"tests"`
}
type HealthCheckTest struct {
	DurationMillis  int64	`json:"duration_millis"`
	TestName        string	`json:"test_name"`
	TestResult      string	`json:"rest_result"`
	TestedAt        string  `json:"tested_at"`	//ISO 8601 Representation
}
