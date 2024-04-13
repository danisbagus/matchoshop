package usecase

import (
	"github.com/danisbagus/matchoshop/internal/repository"
)

type IHealthCheckUsecase interface {
	Get() map[string]interface{}
}

type HealthCheckUsecase struct {
	repo repository.IHealthcheckRepository
}

func NewHealthCheckUsecase(repository repository.RepositoryCollection) IHealthCheckUsecase {
	return &HealthCheckUsecase{
		repo: repository.HealthCheckRepository,
	}
}

func (s HealthCheckUsecase) Get() map[string]interface{} {
	return s.repo.Get()
}
