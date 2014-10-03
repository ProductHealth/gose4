package checks

import (
	"net/url"
	"github.com/ProductHealth/gose4/healthcheck"
	"github.com/ProductHealth/gose4/util"
	"net/http"
	"strings"
	"fmt"
	"io"
)

type httpCheck struct {
	url                       *url.URL
	method                    string
	statusCode int
	requestBody *string
	requestHeaders *map[string]string
	expectedHeaders *map[string]string
	expectedBody *string
	config  healthcheck.HealthCheckConfiguration
}

func NewHttpCheck(address, method string, statusCode int, config healthcheck.HealthCheckConfiguration) (*httpCheck, error){
	url, err:=url.Parse(address)
	if err!= nil {
		return nil, err
	}

	return &httpCheck{url: url, method: method, statusCode: statusCode, config: config }, nil
}

func (hc *httpCheck) RequestBody(content string) {
	hc.requestBody = &content
}

func (hc *httpCheck) RequestHeaders(rh *map[string]string) {
	hc.requestHeaders = rh
}

func (hc *httpCheck) ExpectedHeaders(eh *map[string]string) {
	hc.expectedHeaders = eh
}

func (hc *httpCheck) ExpectedBody(content string) {
	hc.requestBody = &content
}

func (hc *httpCheck) Run() healthcheck.HealthCheckResult {
	sw := util.CreateStopWatch()
	var bodyReader io.Reader = nil
	if hc.requestBody != nil {
		bodyReader = strings.NewReader(*(hc.requestBody))
	}
	request, err := http.NewRequest(hc.method, hc.url.String(), bodyReader)
	if err != nil {
		return healthcheck.Failed(sw.GetDuration(), err.Error())
	}
	client := http.DefaultClient // TODO : Make configurable
	response, err := client.Do(request)
	if hc.statusCode != response.StatusCode {
		return healthcheck.Failed(sw.GetDuration(), fmt.Sprintf("Returned response code %v does not match required %v ", response.StatusCode, hc.statusCode))
	}
	return healthcheck.Ok(sw.GetDuration())
}

func (hc *httpCheck) Configuration() healthcheck.HealthCheckConfiguration {
	return hc.config
}

