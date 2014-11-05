package etcd

import (
	"fmt"
	"github.com/ProductHealth/gose4"
	"github.com/coreos/go-etcd/etcd"
	"os"
	"time"
)

type check struct {
	client *etcd.Client
	config gose4.Configuration
}

func New(client *etcd.Client, config gose4.Configuration) *check {
	return &check{client: client, config: config}
}

func (hc check) Run() gose4.HealthCheckResult {
	sw := gose4.CreateStopWatch()
	hostname, err := os.Hostname()
	path := fmt.Sprintf("/%v/%v", hostname, time.Now().Unix())
	_, err = hc.client.Set(path, "", 1)
	if err != nil {
		return gose4.Failed(sw.GetDuration(), fmt.Sprintf("encountered error while quering etcd : %v", err))
	} else {
		return gose4.Ok(sw.GetDuration())
	}
}

func (hc check) Configuration() gose4.Configuration {
	return hc.config
}
