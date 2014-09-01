package server

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ProductHealth/gose4/healthcheck"
	"github.com/ProductHealth/gose4/util"
	sigar "github.com/cloudfoundry/gosigar"
	"time"
	"fmt"
	"runtime"
)

func createGetServiceStatus(status Status) restful.RouteFunction {
	// Populate static runtime status
	serviceStartTime := time.Now()
	numberOfCpus := runtime.NumCPU()
	status.OsNumberProcessors = &numberOfCpus
	status.MachineName = util.GetCurrentHostName()
	concreteSigar := sigar.ConcreteSigar{}
	status.OsArch = runtime.GOARCH
	status.OsName = runtime.GOOS
	status.OsVersion = "n/a"
	return func(_ *restful.Request, response *restful.Response) {
		currentTime := time.Now()
		res := status
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

func createGetServiceHealthcheck(healthcheckservice *healthcheck.HealthcheckService) restful.RouteFunction {
	return func(request *restful.Request, response *restful.Response) {
		result := HealthCheck{}
		result.ReportAsOf = timeToIso8601(time.Now().UTC())
		result.Tests = []HealthCheckTest{}
		//result.Tests = []healthcheck.HealthCheckResult{}
		for check, lastResult := range healthcheckservice.GetResults() {
			resultItem := HealthCheckTest{}
			resultItem.DurationMillis = lastResult.DurationMillis()
			resultItem.TestName = check.Configuration().Description
			resultItem.TestResult = lastResult.Status.String()
			resultItem.TestedAt = timeToIso8601(lastResult.LastCheck)
			result.Tests = append(result.Tests, resultItem)
		}
		response.WriteEntity(result)
	}
}

func RegisterRestEndpoints(ws *restful.WebService, se4 *healthcheck.HealthcheckService) {
	status := Status{}
	ws.Route(ws.GET("/service/status").To(createGetServiceStatus(status)))
	ws.Route(ws.GET("/service/healthcheck").To(createGetServiceHealthcheck(se4)))
}
func CreateRestServer(service *healthcheck.HealthcheckService) *restful.WebService {
	webService := new(restful.WebService)
	webService.Consumes(restful.MIME_JSON)
	webService.Produces(restful.MIME_JSON)
	webService.Filter(addPoweredByFilter)
	RegisterRestEndpoints(webService, service)
	return webService
}

func addPoweredByFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	response.AddHeader("X-Generated-By", "goSE4")
	chain.ProcessFilter(request, response)
}
