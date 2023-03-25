package service

import (
	"github.com/danisbagus/matchoshop/internal/core/port"
)

type HealthCheckService struct {
	repo port.HealthCheckRepo
}

func NewHealthCheckService(repo port.HealthCheckRepo) port.HealthCheckService {
	return &HealthCheckService{
		repo: repo,
	}
}

func (s HealthCheckService) Get() map[string]interface{} {
	return s.repo.Get()
}
