package http

import (
	"net/url"
	"github.com/ProductHealth/gose4/healthcheck"
	"github.com/ProductHealth/gose4/util"
	"net/http"
	"strings"
	"fmt"
	"io"
)

type HttpCheck struct {
	Uri                       url.URL
	Method                    string
	RequestBody *string
	ExpectedHttpResponse *int
	ExpectedHeader *map[string]string
	ExpectedBodyContent *string
	HealthCheckConfiguration  healthcheck.HealthCheckConfiguration
}

func (httpCheck HttpCheck) Run() healthcheck.HealthCheckResult {
	sw := util.CreateStopWatch()
	var bodyReader io.Reader = nil
	if httpCheck.RequestBody != nil {
		bodyReader = strings.NewReader(*(httpCheck.RequestBody))
	}
	request, err := http.NewRequest(httpCheck.Method, httpCheck.Uri.String(), bodyReader)
	if err != nil {
		return healthcheck.Failed(sw.GetDuration(), err.Error())
	}
	client := http.DefaultClient // TODO : Make configurable
	response, err := client.Do(request)
	if httpCheck.ExpectedHttpResponse != nil && *(httpCheck.ExpectedHttpResponse) != response.StatusCode {
		return healthcheck.Failed(sw.GetDuration(), fmt.Sprintf("Returned response code %v does not match required %v ", response.StatusCode, *httpCheck.ExpectedHttpResponse))
	}
	return healthcheck.Ok(sw.GetDuration())
}

func (hc HttpCheck) Configuration() healthcheck.HealthCheckConfiguration {
	return hc.HealthCheckConfiguration
}

