package checks

import (
	"fmt"
	"github.com/ProductHealth/gose4/healthcheck"
)

func NewDownstreamSE4HttpCheck(address string) (healthcheck.HealthCheck, error) {
	return NewHttpCheck(
		address,
		"GET",
		200,
		healthcheck.NewConfiguration(fmt.Sprintf("Downstream check at %v failed", address), healthcheck.SeverityWarn, 0, 300))
}
