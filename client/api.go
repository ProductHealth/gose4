package client

import (
	"github.com/ProductHealth/gose4"
)

type Client interface {
	Healthcheck() (*gose4.TestResults, error)
}
