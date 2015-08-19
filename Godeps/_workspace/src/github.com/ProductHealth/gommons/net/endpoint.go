package net

import (
	"net/url"
	"strconv"
	"strings"
)

// Description of logical network endpoint
type Endpoint interface {
	HostName() string
	Port() int
}

type defaultEndpoint struct {
	hostname string
	port     int
}

func (e defaultEndpoint) HostName() string {
	return e.hostname
}
func (e defaultEndpoint) Port() int {
	return e.port
}

func NewEndpoint(hostname string, port int) Endpoint {
	return defaultEndpoint{hostname, port}
}

func FromUrl(u url.URL) Endpoint {
	parts := strings.Split(u.Host, ":")
	if len(parts) != 2 {
		return nil
	}
	port, _ := strconv.Atoi(parts[1])
	return NewEndpoint(parts[0], port)
}
