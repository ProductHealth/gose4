package http

import (
	"fmt"
	"github.com/ProductHealth/gose4"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/ProductHealth/gommons/time"
)

type check struct {
	url             *url.URL
	method          string
	statusCode      int
	requestBody     *string
	requestHeaders  *map[string]string
	expectedHeaders *map[string]string
	expectedBody    *string
	config          gose4.Configuration
	requestFunc     func(*http.Request) (*http.Response, error)
}

func New(address, method string, statusCode int, config gose4.Configuration) (*check, error) {
	url, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	requestFunc := func(req *http.Request) (*http.Response, error) {
		return http.DefaultClient.Do(req)
	}

	return &check{url: url, method: method, statusCode: statusCode, config: config, requestFunc: requestFunc}, nil
}

func (hc *check) RequestBody(content string) {
	hc.requestBody = &content
}

func (hc *check) RequestHeaders(rh map[string]string) {
	hc.requestHeaders = &rh
}

func (hc *check) ExpectedHeaders(eh map[string]string) {
	hc.expectedHeaders = &eh
}

func (hc *check) ExpectedBody(content string) {
	hc.requestBody = &content
}

func (hc *check) Run() gose4.HealthCheckResult {
	sw := time.NewStopWatch()
	var bodyReader io.Reader = nil
	if hc.requestBody != nil {
		bodyReader = strings.NewReader(*(hc.requestBody))
	}
	request, err := http.NewRequest(hc.method, hc.url.String(), bodyReader)
	if err != nil {
		return gose4.Failed(sw.GetDuration(), err.Error())
	}
	response, err := hc.requestFunc(request)
	gose4.Debugf("RESPONSE %#v ERR: %s", response, err)
	if err != nil {
		return gose4.Failed(sw.GetDuration(), fmt.Sprintf("Error while requesting %#v: %s", hc.url, err))
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	if hc.statusCode != response.StatusCode {
		return gose4.Failed(sw.GetDuration(), fmt.Sprintf("Returned response code %#v does not match required %#v ", response.StatusCode, hc.statusCode))
	}
	return gose4.Ok(sw.GetDuration())
}

func (hc *check) Configuration() gose4.Configuration {
	return hc.config
}
