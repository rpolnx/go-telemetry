package service

import "github.com/samber/do"

type HealthCheckService interface {
	Check() string
}

type healthCheckService struct {
}

func (healthCheckService) Check() string {
	return "OK"
}

func NewHealthCheckService(i *do.Injector) (HealthCheckService, error) {
	return &healthCheckService{}, nil
}
