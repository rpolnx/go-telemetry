package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rpolnx/go-telemetry/internal/config"
	"github.com/rpolnx/go-telemetry/internal/server"
	"github.com/rpolnx/go-telemetry/internal/service"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
)

type HealthCheckController interface {
	BaseController
	Get(c *gin.Context)
}

type healthCheckController struct {
	svc    service.HealthCheckService
	logger *logrus.Logger
}

func (c *healthCheckController) RegisterRoutes(e *gin.Engine) {
	e.GET("/healthcheck", c.Get)
}

func (hc healthCheckController) Get(c *gin.Context) {
	c.JSON(200, gin.H{"status": hc.svc.Check()})
}

func NewHealthCheckController(ioc *do.Injector) (HealthCheckController, error) {
	healthCheckService := do.MustInvoke[service.HealthCheckService](ioc)
	_ = do.MustInvoke[*config.Config](ioc)
	logger := do.MustInvoke[*logrus.Logger](ioc)
	engine := do.MustInvoke[*server.Server](ioc)

	hc := &healthCheckController{healthCheckService, logger}
	hc.RegisterRoutes(engine.HttpServer)

	return hc, nil
}
