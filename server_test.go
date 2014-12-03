package gose4

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerFunc(t *testing.T) {
	healthCheckService := New()
	handler := HandlerFunc(healthCheckService)

	req, err := http.NewRequest("GET", "http://example.com/service/status", nil)
	if err != nil {
		t.Errorf("Did not expect an error: %s", err)
	}

	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200 but go: %d", w.Code)
	}
}

func TestGoodToGoPassesWhenNoTestsPresent(t *testing.T) {
	healthCheckService := New()
	handler := HandlerFunc(healthCheckService)

	req, err := http.NewRequest("GET", "http://example.com/service/healthcheck/gtg", nil)
	if err != nil {
		t.Errorf("Did not expect an error: %s", err)
	}

	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200 but go: %d", w.Code)
	}
}

func TestGoodToGoFailsWhenSeverityNoticeFails(t *testing.T) {
	healthCheckService := New()
	handler := HandlerFunc(healthCheckService)
	check := NewTestCheck()
	check.TestCheckResult.Result = CheckFailed
	check.TestCheckConfig.Severity = SeverityNotice
	healthCheckService.Register(check)
	healthCheckService.executeHealthCheck(check)
	req, err := http.NewRequest("GET", "http://example.com/service/healthcheck/gtg", nil)
	if err != nil {
		t.Errorf("Did not expect an error: %s", err)
	}

	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200 but go: %d", w.Code)
	}
}

func TestGoodToGoPassesWhenSeverityFatalFails(t *testing.T) {
	healthCheckService := New()
	handler := HandlerFunc(healthCheckService)
	check := NewTestCheck()
	check.TestCheckResult.Result = CheckFailed
	healthCheckService.Register(check)
	healthCheckService.executeHealthCheck(check)
	req, err := http.NewRequest("GET", "http://example.com/service/healthcheck/gtg", nil)
	if err != nil {
		t.Errorf("Did not expect an error: %s", err)
	}

	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != 503 {
		t.Errorf("Expected 503 but go: %d", w.Code)
	}
}
