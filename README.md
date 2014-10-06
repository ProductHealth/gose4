# Product Health gose4
Product Health gose4 is a golang implementation of Beamly [se4](https://github.com/beamly/SE4) library to provide
a simple convention for getting standard service information.

Gose4 standard service information is delivered through an HTTP interface.  It can spawn its own web server on a given
port or integrate to an existing one through the [HandleFunc](http://golang.org/pkg/net/http/#HandlerFunc) interface.

# Example

    package main

    import (
    	"github.com/ProductHealth/gose4"
    	"github.com/ProductHealth/gose4/http"
    )

    func main() {
    	healthCheckService := gose4.New()
    	check, err := http.New(
    		"http://example.com",
    		"GET",
    		200,
    		gose4.NewConfiguration("Example Check", gose4.SeverityWarn, 0, 5))

    	if err != nil {
    		glog.Fatal(err)
    	}

    	healthCheckService.Register(check)

    	go gose4.StartHttpServer(healthCheckService, 8081)
    }

## Service Status

GET /service/status will return a json document like the following:

    {
      "artifact_id": "kats",
      "build_number": "dev",
      "build_machine": "emeka.local",
      "build_by": "ph",
      "build_when": "2014-10-06T12:05:29Z",
      "compiler_version": "go1.3",
      "current_time": "2014-10-06T12:05:48Z",
      "git_sha": "0e2b1d833317d70e50e92c633828f93a7f57a598",
      "machine_name": "emeka.local",
      "os_arch": "amd64",
      "os_name": "darwin",
      "os_version": "n/a",
      "runbook_uri": "",
      "up_duration": "13.955480359 seconds",
      "up_since": "2014-10-06T13:05:34Z",
      "version": "",
      "os_avgload": "1.83544921875",
      "os_numprocessors": 8
     }

## Health Check

GET /service/healthcheck will return as json document like the following:

    {
      "report_as_of": "2014-10-06T14:37:20Z",
      "report_duration": "",
      "tests": [
       {
        "duration_millis": 3,
        "test_name": "Kairosdb health",
        "rest_result": "passed",
        "tested_at": "2014-10-06T14:37:19Z"
       }
      ]
     }

where zero or more health check results will be returned.