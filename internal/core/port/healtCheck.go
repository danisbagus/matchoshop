package port

type HealthCheckService interface {
	Get() map[string]interface{}
}
