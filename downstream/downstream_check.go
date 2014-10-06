package downstream

import (
	"fmt"
	"github.com/ProductHealth/gose4"
	"github.com/ProductHealth/gose4/http"
)

func NewCheck(address string) (gose4.HealthCheck, error) {
	return http.NewCheck(
		address,
		"GET",
		200,
		gose4.NewConfiguration(fmt.Sprintf("Downstream check %v", address), gose4.SeverityWarn, 0, 300))
}
