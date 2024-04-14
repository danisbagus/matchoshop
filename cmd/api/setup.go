package api

import (
	"github.com/danisbagus/matchoshop/cmd/middleware"
	handlerV1 "github.com/danisbagus/matchoshop/internal/handler/http/v1"
	"github.com/danisbagus/matchoshop/internal/usecase"
	"github.com/gorilla/mux"
)

func Set(r *mux.Router, usecaseCollection usecase.UsecaseCollection, APIMiddleware middleware.IAPIMiddleware) {
	handlerV1.NewHealthcheckHandler(r, usecaseCollection, APIMiddleware)
	handlerV1.NewOrderHandler(r, usecaseCollection, APIMiddleware)
	handlerV1.NewProductCategoryHandler(r, usecaseCollection, APIMiddleware)
	handlerV1.NewProductHandler(r, usecaseCollection, APIMiddleware)
	handlerV1.NewReviewHandler(r, usecaseCollection, APIMiddleware)
	handlerV1.NewUploadHandler(r, usecaseCollection, APIMiddleware)
	handlerV1.NewUserHandler(r, usecaseCollection, APIMiddleware)
}
