package checks

import (
	"github.com/ProductHealth/gose4/healthcheck"
	"testing"
	"net/http"
	"net/url"
)

func TestRun(t *testing.T){
	requestFunc:= func (req *http.Request) (*http.Response,error) {
		return &http.Response{StatusCode: 200, Request: req}, nil
	}
	url, _:=url.Parse("http://foo.bar")

	check:= httpCheck{url:url, method: "GET", statusCode: 200, requestFunc: requestFunc}

	result:=check.Run()

	if result.Status != healthcheck.StatusPassed {
		t.Errorf("Expected passed but got: %#v", result)
	}
}
