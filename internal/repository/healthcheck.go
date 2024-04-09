package repository

import (
	"github.com/jmoiron/sqlx"
)

type IHealthcheckRepository interface {
	Get() (result map[string]interface{})
}

type HealthCheckRepo struct {
	db *sqlx.DB
}

func NewHealthCheckRepository(db *sqlx.DB) *HealthCheckRepo {
	return &HealthCheckRepo{
		db: db,
	}
}
func (r HealthCheckRepo) Get() (result map[string]interface{}) {
	result = map[string]interface{}{"ping": false}
	if err := r.db.Ping(); err != nil {
		return
	}

	result["ping"] = true
	return result
}
