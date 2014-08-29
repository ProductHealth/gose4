package server

import (
	restful "github.com/emicklei/go-restful"
	"github.com/ProductHealth/gose4/healthcheck"
	"time"
)

func createGetServiceStatus(healthcheckservice *healthcheck.HealthcheckService) restful.RouteFunction {
	return func(request *restful.Request, response *restful.Response) {
		healthcheckservice.GetResults()
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
	ws.Route(ws.GET("/service/status").To(createGetServiceStatus(se4)))
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

func timeToIso8601(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z")
}

func addPoweredByFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	response.AddHeader("X-Generated-By", "goSE4")
	chain.ProcessFilter(request, response)
}
