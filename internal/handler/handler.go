package handler

import (
	"fmt"

	"github.com/rpolnx/go-telemetry/internal/config"
	"github.com/rpolnx/go-telemetry/internal/controller"
	"github.com/rpolnx/go-telemetry/internal/server"
	"github.com/rpolnx/go-telemetry/internal/service"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
)

func GetInjectorManager() *do.Injector {
	ioc := do.NewWithOpts(&do.InjectorOpts{
		HookAfterRegistration: func(injector *do.Injector, serviceName string) {
			fmt.Printf("Service registered: %s\n", serviceName)
		},
		HookAfterShutdown: func(injector *do.Injector, serviceName string) {
			fmt.Printf("Service stopped: %s\n", serviceName)
		},

		Logf: func(format string, args ...any) {
			logrus.Infof(format, args...)
		},
	})

	ProvideGeneralSetup(ioc)
	ProvideRepositories(ioc)
	ProvideServices(ioc)
	ProvideControllers(ioc)

	return ioc
}

func ProvideGeneralSetup(ioc *do.Injector) {
	do.Provide(ioc, config.NewConfig)
	do.Provide(ioc, config.NewLogger)
	do.Provide(ioc, server.NewServer)
}

func ProvideRepositories(ioc *do.Injector) {

}

func ProvideServices(ioc *do.Injector) {
	do.Provide(ioc, service.NewHealthCheckService)
}

func ProvideControllers(ioc *do.Injector) {
	do.Provide(ioc, controller.NewHealthCheckController)

	//self initializing
	do.MustInvoke[controller.HealthCheckController](ioc)
}
