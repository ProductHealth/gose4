package gose4

import (
	"fmt"
	sigar "github.com/cloudfoundry/gosigar"
	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
	"runtime"
	"time"
)

type TestResults struct {
	ReportAsOf     string            `json:"report_as_of"`    //ISO 8601 Representation
	ReportDuration string            `json:"report_duration"` //
	Tests          []TestResult `json:"tests"`
}
type TestResult struct {
	DurationMillis int64  `json:"duration_millis"`
	TestName       string `json:"test_name"`
	TestResult     string `json:"rest_result"`
	TestedAt       string `json:"tested_at"` //ISO 8601 Representation
}

func StartHttpServer(service *healthCheckService, httpPort int) {
	container := restful.NewContainer()
	glog.Infof("Starting SE4 server on port %v", httpPort)
	container.Add(createRestServer(service))
	httpServer := &http.Server{Addr: fmt.Sprintf(":%v", httpPort), Handler: container}

	httpServer.ListenAndServe()
}

func HandlerFunc(service *healthCheckService) http.HandlerFunc {
	container := restful.NewContainer()
	container.Add(createRestServer(service))
	return func(w http.ResponseWriter, r *http.Request) {
		container.ServeHTTP(w, r)
	}
}

func createGetServiceStatus() restful.RouteFunction {
	// Populate static runtime status
	serviceStartTime := time.Now()
	numberOfCpus := runtime.NumCPU()
	ServiceStatus.OsNumberProcessors = &numberOfCpus
	ServiceStatus.MachineName = GetCurrentHostName()
	concreteSigar := sigar.ConcreteSigar{}
	ServiceStatus.OsArch = runtime.GOARCH
	ServiceStatus.OsName = runtime.GOOS
	ServiceStatus.OsVersion = "n/a"
	return func(_ *restful.Request, response *restful.Response) {
		currentTime := time.Now()
		res := ServiceStatus
		// Time related field
		res.SetCurrentTime(&currentTime)
		res.UpSince = timeToIso8601(serviceStartTime)
		uptime := currentTime.Sub(serviceStartTime)
		res.UpDuration = fmt.Sprintf("%v seconds", uptime.Seconds())
		// Get load avg
		loadAvg, _ := concreteSigar.GetLoadAverage()
		loadAvgString := fmt.Sprintf("%v", loadAvg.Five)
		res.OsLoad = &loadAvgString

		response.WriteEntity(res)
	}
}

func createGetServiceHealthCheck(healthcheckservice *healthCheckService) restful.RouteFunction {
	return func(request *restful.Request, response *restful.Response) {
		result := TestResults{}
		result.ReportAsOf = timeToIso8601(time.Now().UTC())
		result.Tests = []TestResult{}
		for check, lastResult := range healthcheckservice.GetResults() {
			resultItem := TestResult{}
			resultItem.DurationMillis = lastResult.DurationMillis()
			resultItem.TestName = check.Configuration().Description
			resultItem.TestResult = lastResult.Result.String()
			resultItem.TestedAt = timeToIso8601(lastResult.LastCheck)
			result.Tests = append(result.Tests, resultItem)
		}
		response.WriteEntity(result)
	}
}

func registerRestEndpoints(ws *restful.WebService, se4 *healthCheckService) {
	ws.Route(ws.GET("/service/status").To(createGetServiceStatus()))
	ws.Route(ws.GET("/service/healthcheck").To(createGetServiceHealthCheck(se4)))
}
func createRestServer(service *healthCheckService) *restful.WebService {
	webService := new(restful.WebService)
	webService.Consumes(restful.MIME_JSON)
	webService.Produces(restful.MIME_JSON)
	webService.Filter(addPoweredByFilter)
	registerRestEndpoints(webService, service)
	return webService
}

func addPoweredByFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	response.AddHeader("X-Generated-By", "goSE4")
	chain.ProcessFilter(request, response)
}

func timeToIso8601(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z")
}
