package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rpolnx/go-telemetry/internal/config"
	"github.com/rpolnx/go-telemetry/internal/handler"
	"github.com/rpolnx/go-telemetry/internal/server"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
)

func main() {
	ioc := handler.GetInjectorManager()

	cfg := do.MustInvoke[*config.Config](ioc)
	logger := do.MustInvoke[*logrus.Logger](ioc)

	cleanup := config.InitTracer()
	defer cleanup(context.Background())

	engine := do.MustInvoke[*server.Server](ioc)

	go engine.HttpServer.Run(fmt.Sprintf(":%d", cfg.Port))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	ioc.ShutdownOnSignals(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	logger.Infof("Server being finalized")
}
