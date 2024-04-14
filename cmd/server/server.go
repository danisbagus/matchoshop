package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/cmd/api"
	"github.com/danisbagus/matchoshop/cmd/middleware"
	"github.com/danisbagus/matchoshop/infrastructure/config"
	"github.com/danisbagus/matchoshop/infrastructure/database"
	"github.com/danisbagus/matchoshop/internal/repository"
	"github.com/danisbagus/matchoshop/internal/usecase"
	"github.com/gorilla/mux"
)

type Server struct {
	Http *httpServer
}

type httpServer struct {
}

func NewServer() *Server {
	return &Server{
		Http: &httpServer{},
	}
}

func (s *httpServer) Start() {
	config.SetConfig(".", ".env")

	dbConnection := database.CreateDBConnections(database.GetConfigs())

	r := mux.NewRouter()

	repositoryCollection := repository.NewRepositoryCollection(dbConnection.Postgres)
	usecaseCollection := usecase.NewUsecaseCollection(repositoryCollection)
	APIMiddleware := middleware.NewAPIMiddleware()

	middleware.Set(r)
	api.Set(r, usecaseCollection, APIMiddleware)

	appPort := fmt.Sprintf("%s:%v", config.APP_HOST, config.APP_PORT)

	err := http.ListenAndServe(appPort, r)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info(fmt.Sprintf("Starting the application at:%s", appPort))
}
