package v1

import (
	"net/http"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/matchoshop/cmd/middleware"
	"github.com/danisbagus/matchoshop/internal/usecase"
	"github.com/gorilla/mux"
)

type UploadHandler struct {
	uploadUsecase usecase.IUploadUsecase
}

func NewUploadHandler(r *mux.Router, usecaseCollection usecase.UsecaseCollection, APIMiddleware middleware.IAPIMiddleware) {
	handler := UploadHandler{
		uploadUsecase: usecaseCollection.UploadUsecase,
	}

	route := r.PathPrefix("/api/v1/upload").Subrouter()
	route.HandleFunc("/image", handler.UploadImage).Methods(http.MethodPost)
}

func (h UploadHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	//  Ensure our file does not exceed 5MB
	r.Body = http.MaxBytesReader(w, r.Body, 5*1024*1024)

	file, _, err := r.FormFile("file")
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	defer file.Close()

	url, appErr := h.uploadUsecase.UploadImage(file)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	resData := map[string]interface{}{
		"message": "Successfully upload image",
		"url":     url,
	}
	response.Write(w, http.StatusOK, resData)
}
