package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rpolnx/go-telemetry/internal/config"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"

	ginlogrus "github.com/toorop/gin-logrus"
)

type Server struct {
	HttpServer *gin.Engine
}

func NewServer(ioc *do.Injector) (*Server, error) {
	_ = do.MustInvoke[*config.Config](ioc)
	logger := do.MustInvoke[*logrus.Logger](ioc)

	r := gin.New()

	r.Use(ginlogrus.Logger(logger), gin.Recovery())

	return &Server{
		HttpServer: r,
	}, nil
}
