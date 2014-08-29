package http

import (
	"net/url"
	"github.com/ProductHealth/gose4/healthcheck"
	"github.com/ProductHealth/gose4/util"
	"net/http"
	"strings"
	"fmt"
)

type HttpCheck struct {
	Uri         url.URL
	Method      string
	RequestBody *string
	ExpectedHttpResponse *int
	ExpectedHeader *map[string]string
	ExpectedBodyContent *string
}

func (httpCheck HttpCheck) Run() HealthCheckResult {
	sw := util.CreateStopWatch()
	var bodyReader io.Reader = nil
	if httpCheck.RequestBody != nil {
		bodyReader = strings.NewReader(httpCheck.RequestBody)
	}
	request, err := http.NewRequest(httpCheck.Method, httpCheck.Uri.String(), bodyReader)
	if err != nil {
		return HealthCheckResult{sw.GetDuration(), healthcheck.Failed, err.Error()}
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if httpCheck.ExpectedHttpResponse != nil && httpCheck.ExpectedHttpResponse != response.StatusCode {
		return HealthCheckResult{sw.GetDuration(), healthcheck.Failed, fmt.Fprintf("Returned response code %v does not match required %v ", response.StatusCode, httpCheck.ExpectedHttpResponse)}
	}
	return HealthCheckResult{sw.GetDuration(), healthcheck.Passed, "OK"}
}
