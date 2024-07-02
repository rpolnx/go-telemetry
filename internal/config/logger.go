package config

import (
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
)

func NewLogger(i *do.Injector) (*logrus.Logger, error) {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)

	return logger, nil
}
