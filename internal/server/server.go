package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rpolnx/go-telemetry/internal/config"
	"github.com/rpolnx/go-telemetry/internal/controller"
	"github.com/rpolnx/go-telemetry/internal/service"
	"github.com/samber/do"
)

func GetServer() {

	ioc := do.New()

	do.Provide(ioc, config.NewConfig)

	do.Provide(ioc, service.NewHealthCheckService)
	do.Provide(ioc, controller.NewHealthCheckController)

	do.MustInvoke[controller.HealthCheckController](ioc)

	cfg := do.MustInvoke[*config.Config](ioc)

	fmt.Printf("Starting server at port %d", cfg.Port)

	go http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	fmt.Println("killing")
}
