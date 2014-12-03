package client

import (
	gnet "github.com/ProductHealth/gommons/net"
	"github.com/ProductHealth/gose4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClientParsesSE4Result(t *testing.T) {
	testCheck := gose4.NewTestCheck()
	se4 := gose4.New()
	se4.Register(testCheck)
	server := httptest.NewServer(gose4.HandlerFunc(se4))
	defer server.Close()
	u, _ := url.Parse(server.URL)
	client := New(gnet.FromUrl(*u))
	result, err := client.Healthcheck()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testCheck.TestCheckResult.DurationMillis(), result.Tests[0].DurationMillis)
	assert.Equal(t, testCheck.TestCheckResult.Result.String(), result.Tests[0].TestResult)
	assert.Equal(t, testCheck.TestCheckResult.LastCheck.Format(gose4.SE4TimeFormat), result.Tests[0].TestedAt)
	assert.Equal(t, testCheck.TestCheckConfig.Description, result.Tests[0].TestName)
}
