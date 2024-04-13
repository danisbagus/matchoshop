package v1

import (
	"net/http"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/matchoshop/internal/usecase"
)

type HealthCheckHandler struct {
	usecase usecase.IHealthCheckUsecase
}

func NewHealthCheckHandlerHandler(usecase usecase.IHealthCheckUsecase) *HealthCheckHandler {
	return &HealthCheckHandler{
		usecase: usecase,
	}
}

func (h HealthCheckHandler) Get(w http.ResponseWriter, r *http.Request) {
	get := h.usecase.Get()

	resData := map[string]interface{}{
		"message": "Successfully get health-check",
		"data":    get,
	}
	response.Write(w, http.StatusOK, resData)
}
