package v1

import (
	"net/http"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/matchoshop/internal/core/port"
)

type HealthCheckHandler struct {
	service port.HealthCheckService
}

func NewHealthCheckHandlerHandler(service port.HealthCheckService) *HealthCheckHandler {
	return &HealthCheckHandler{
		service: service,
	}
}

func (h HealthCheckHandler) Get(w http.ResponseWriter, r *http.Request) {
	get := h.service.Get()

	resData := map[string]interface{}{
		"message": "Successfully get health-check",
		"data":    get,
	}
	response.Write(w, http.StatusOK, resData)
}
