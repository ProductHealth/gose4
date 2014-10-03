package checks

import (
	"fmt"
	"github.com/ProductHealth/gose4/healthcheck"
	"github.com/ProductHealth/gose4/util"
	"github.com/gocql/gocql"
)

type CassandraConnectionCheck struct {
	Session                  *gocql.Session
	HealthCheckConfiguration healthcheck.HealthCheckConfiguration
}

func (hc CassandraConnectionCheck) Run() healthcheck.HealthCheckResult {
	sw := util.CreateStopWatch()
	valid, err := hc.checkConnection()
	if !valid {
		return healthcheck.Failed(sw.GetDuration(), fmt.Sprintf("encountered error while quering cassandra : %v", err))
	} else {
		return healthcheck.Ok(sw.GetDuration())
	}
}

func (hc CassandraConnectionCheck) checkConnection() (bool, error) {
	res, err := hc.Session.Query("SELECT now() FROM system.local").Iter().RowData()
	return len(res.Values) == 1, err
}

func (hc CassandraConnectionCheck) Configuration() healthcheck.HealthCheckConfiguration {
	return hc.HealthCheckConfiguration
}
