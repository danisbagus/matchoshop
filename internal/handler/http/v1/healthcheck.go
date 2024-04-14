package v1

import (
	"net/http"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/matchoshop/cmd/middleware"
	"github.com/danisbagus/matchoshop/internal/usecase"
	"github.com/gorilla/mux"
)

type HealthcheckHandler struct {
	healthcheckUsecase usecase.IHealthcheckUsecase
}

func NewHealthcheckHandler(r *mux.Router, usecaseCollection usecase.UsecaseCollection, APIMiddleware middleware.IAPIMiddleware) {
	handler := HealthcheckHandler{
		healthcheckUsecase: usecaseCollection.HealthcheckUsecase,
	}

	route := r.PathPrefix("/api/v1/health-check").Subrouter()
	route.HandleFunc("", handler.Get).Methods(http.MethodGet)

}

func (h HealthcheckHandler) Get(w http.ResponseWriter, r *http.Request) {
	get := h.healthcheckUsecase.Get()

	resData := map[string]interface{}{
		"message": "Successfully get health-check",
		"data":    get,
	}
	response.Write(w, http.StatusOK, resData)
}
