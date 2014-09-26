package server

import (
	"github.com/ProductHealth/gose4/healthcheck"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerFunc(t *testing.T) {
	healthCheckService := healthcheck.CreateHealthcheckService()
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
