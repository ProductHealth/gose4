package checks

import (
	"github.com/ProductHealth/gose4/healthcheck"
	"fmt"
)

func NewDownstreamSE4HttpCheck(address string) (healthcheck.HealthCheck,error) {
	return NewHttpCheck(
		address,
		"GET",
		200,
		healthcheck.NewDefaultConfiguration(fmt.Sprintf("Downstream check at %v failed", address), healthcheck.SeverityWarn))
}
