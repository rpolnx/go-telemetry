package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rpolnx/go-telemetry/internal/config"
	"github.com/rpolnx/go-telemetry/internal/service"
	"github.com/samber/do"
)

type HealthCheckController interface {
	BaseController
	Get(w http.ResponseWriter, r *http.Request)
}

type healthCheckController struct {
	svc service.HealthCheckService
}

func (c *healthCheckController) RegisterRoutes() {
	http.HandleFunc("/healthcheck", c.Get)
}

func (c healthCheckController) Get(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": c.svc.Check()}
	resp, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func NewHealthCheckController(ioc *do.Injector) (HealthCheckController, error) {
	healthCheckService := do.MustInvoke[service.HealthCheckService](ioc)
	_ = do.MustInvoke[*config.Config](ioc)

	hc := &healthCheckController{healthCheckService}
	hc.RegisterRoutes()

	return hc, nil
}
