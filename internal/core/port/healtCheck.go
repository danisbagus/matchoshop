package port

type HealthCheckRepo interface {
	Get() map[string]interface{}
}

type HealthCheckService interface {
	Get() map[string]interface{}
}
