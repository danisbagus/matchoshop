package v1

import (
	"net/http"

	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/labstack/echo/v4"
)

type HealthCheckHandler struct {
	service port.HealthCheckService
}

func NewHealthCheckHandlerHandler(service port.HealthCheckService) *HealthCheckHandler {
	return &HealthCheckHandler{
		service: service,
	}
}

func (h HealthCheckHandler) Get(c echo.Context) error {
	get := h.service.Get()
	resData := map[string]interface{}{
		"message": "Successfully get health-check",
		"data":    get,
	}

	return c.JSON(http.StatusOK, resData)
}
