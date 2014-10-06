package gose4

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerFunc(t *testing.T) {
	healthCheckService := CreateHealthcheckService()
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
