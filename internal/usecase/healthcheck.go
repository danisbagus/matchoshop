package usecase

import (
	"github.com/danisbagus/matchoshop/internal/repository"
)

type IHealthcheckUsecase interface {
	Get() map[string]interface{}
}

type HealthcheckUsecase struct {
	repo repository.IHealthcheckRepository
}

func NewHealthcheckUsecase(repository repository.RepositoryCollection) IHealthcheckUsecase {
	return &HealthcheckUsecase{
		repo: repository.HealthCheckRepository,
	}
}

func (s HealthcheckUsecase) Get() map[string]interface{} {
	return s.repo.Get()
}
