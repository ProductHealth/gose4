package cassandra

import (
	"fmt"
	"github.com/ProductHealth/gose4"
	"github.com/gocql/gocql"
)

type check struct {
	session                  *gocql.Session
	config gose4.Configuration
}

func New(session *gocql.Session, config gose4.Configuration) *check{
	return &check{session: session, config: config}
}

func (hc check) Run() gose4.HealthCheckResult {
	sw := gose4.CreateStopWatch()
	valid, err := hc.checkConnection()
	if !valid {
		return gose4.Failed(sw.GetDuration(), fmt.Sprintf("encountered error while quering cassandra : %v", err))
	} else {
		return gose4.Ok(sw.GetDuration())
	}
}

func (hc check) checkConnection() (bool, error) {
	res, err := hc.session.Query("SELECT now() FROM system.local").Iter().RowData()
	return len(res.Values) == 1, err
}

func (hc check) Configuration() gose4.Configuration {
	return hc.config
}
