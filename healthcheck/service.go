package healthcheck

import (
	"github.com/golang/glog"
	"time"
)

type HealthcheckService struct {
	healthchecks []HealthCheck
	results      map[HealthCheck]HealthCheckResult
}

func CreateHealthcheckService() *HealthcheckService {
	return &HealthcheckService{[]HealthCheck{}, make(map[HealthCheck]HealthCheckResult)}
}

func (se4server HealthcheckService) GetResults() map[HealthCheck]HealthCheckResult {
	return se4server.results
}

func (se4service HealthcheckService) RegisterHealthcheck(h HealthCheck) {
	glog.V(0).Infof("Registering healthcheck '%v'", h.Configuration().Description)
	se4service.healthchecks = append(se4service.healthchecks, h)
	go se4service.runHealthcheck(h)
}

func (se4server HealthcheckService) runHealthcheck(h HealthCheck) {
	// Wait for initial delay to pass
	glog.V(1).Infof("Waiting %v before executing %v", h.Configuration().InitialDelay, h.Configuration().Description)
	time.Sleep(h.Configuration().InitialDelay)
	for {
		glog.V(0).Infof("Executing healthcheck %v", h.Configuration().Description)
		result := h.Run()
		glog.V(1).Infof("Healthcheck returned result : %v", result)
		se4server.results[h] = result
		time.Sleep(h.Configuration().RunDelay)
	}
}
