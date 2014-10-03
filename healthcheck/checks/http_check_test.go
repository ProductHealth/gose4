package checks

import (
	"github.com/ProductHealth/gose4/healthcheck"
	"testing"
	"net/http"
	"net/url"
	"net/http/httptest"
	"fmt"
	"time"
)

func TestRun(t *testing.T) {
	requestFunc := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Request: req}, nil
	}
	url, _ := url.Parse("http://foo.bar")

	check := httpCheck{url:url, method: "GET", statusCode: 200, requestFunc: requestFunc}

	result := check.Run()

	if result.Status != healthcheck.StatusPassed {
		t.Errorf("Expected passed but got: %#v", result)
	}
}

func TestIntegration(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()


	healthCheckService := healthcheck.CreateHealthcheckService()
	checkKairosdb, err := NewHttpCheck(
		ts.URL,
		"GET",
		200,
		healthcheck.NewDefaultConfiguration("Kairosdb not responding", healthcheck.SeverityWarn))

	if err != nil {
		t.Errorf("Did not expect an error: %s", err)
	}

	healthCheckService.RegisterHealthcheck(checkKairosdb)

	var results map[healthcheck.HealthCheck]healthcheck.HealthCheckResult
	for {
		results = healthCheckService.GetResults()
		if len(results) > 0 {
			break
		}
		time.Sleep(time.Millisecond)
	}

	if len(results) != 1 {
		t.Errorf("Expected 1 result but got: %d", len(results))
	}

	if results[checkKairosdb].Status != healthcheck.StatusPassed {
		t.Errorf("Expected StatusPassed but got: %#v", results[checkKairosdb].Status)
	}
}
