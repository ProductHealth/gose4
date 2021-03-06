package http

import (
	"fmt"
	"github.com/ProductHealth/gose4"
	"net/http"
	"net/url"
	"testing"
)

func TestRun(t *testing.T) {
	requestFunc := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Request: req}, nil
	}
	url, _ := url.Parse("http://foo.bar")

	check := check{url: url, method: "GET", statusCode: 200, requestFunc: requestFunc}

	result := check.Run()

	if result.Result != gose4.CheckPassed {
		t.Errorf("Expected passed but got: %#v", result)
	}
}

func TestRunError(t *testing.T) {
	requestFunc := func(req *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("An error")
	}
	url, _ := url.Parse("http://foo.bar")

	check := check{url: url, method: "GET", statusCode: 200, requestFunc: requestFunc}

	result := check.Run()

	if result.Result != gose4.CheckFailed {
		t.Errorf("Expected passed but got: %#v", result)
	}
}
