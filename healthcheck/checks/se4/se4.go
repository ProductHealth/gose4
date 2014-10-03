package se4

import (
	"github.com/ProductHealth/gose4/healthcheck/checks/http"
	"net/url"
)

func NewDownstreamSE4HttpCheck(baseUrl url.URL) http.HttpCheck {
	return http.HttpCheck{}
}
