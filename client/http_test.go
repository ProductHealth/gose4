package client

import (
	gnet "github.com/ProductHealth/gommons/net"
	"github.com/ProductHealth/gose4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

type TestCheck struct{}

var testCheckResult = gose4.HealthCheckResult{
	Duration:  time.Second * 5,
	Result:    gose4.CheckPassed,
	Message:   "It works!",
	LastCheck: time.Unix(1415030870, 0),
}
var testCheckConfig = gose4.Configuration{
	Severity:     gose4.SeverityWarn,
	InitialDelay: time.Second * 0,
	RunDelay:     time.Second,
	Description:  "Example config",
}

func (t TestCheck) Run() gose4.HealthCheckResult {
	return testCheckResult
}
func (t TestCheck) Configuration() gose4.Configuration {
	return testCheckConfig
}

func TestClientParsesSE4Result(t *testing.T) {
	testCheck := TestCheck{}
	se4 := gose4.New()
	se4.Register(testCheck)
	server := httptest.NewServer(gose4.HandlerFunc(se4))
	defer server.Close()
	u, _ := url.Parse(server.URL)
	client := New(gnet.FromUrl(*u))
	result, err := client.Healthcheck()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testCheckResult.DurationMillis(), result.Tests[0].DurationMillis)
	assert.Equal(t, testCheckResult.Result.String(), result.Tests[0].TestResult)
	assert.Equal(t, testCheckResult.LastCheck.Format(gose4.SE4TimeFormat), result.Tests[0].TestedAt)
	assert.Equal(t, testCheckConfig.Description, result.Tests[0].TestName)
}
