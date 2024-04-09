package service

import (
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/repository"
)

type HealthCheckService struct {
	repo repository.IHealthcheckRepository
}

func NewHealthCheckService(repository repository.RepositoryCollection) port.HealthCheckService {
	return &HealthCheckService{
		repo: repository.HealthCheckRepository,
	}
}

func (s HealthCheckService) Get() map[string]interface{} {
	return s.repo.Get()
}
