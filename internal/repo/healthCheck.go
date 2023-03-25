package repo

import (
	"github.com/jmoiron/sqlx"
)

type HealthCheck struct {
	db *sqlx.DB
}

func NewHealthCheck(db *sqlx.DB) *HealthCheck {
	return &HealthCheck{
		db: db,
	}
}
func (r HealthCheck) Get() (result map[string]interface{}) {
	result = map[string]interface{}{"ping": false}
	if err := r.db.Ping(); err != nil {
		return
	}

	result["ping"] = true
	return result
}
