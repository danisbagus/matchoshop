package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danisbagus/matchoshop/cmd/api"
	"github.com/danisbagus/matchoshop/cmd/middleware"
	"github.com/danisbagus/matchoshop/internal/repository"
	"github.com/danisbagus/matchoshop/internal/usecase"
	"github.com/danisbagus/matchoshop/utils/modules"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed loading .env file")
	}

	client := modules.GetPostgresClient()

	defer client.Close()

	r := mux.NewRouter()

	repositoryCollection := repository.NewRepositoryCollection(client)
	usecaseCollection := usecase.NewUsecaseCollection(repositoryCollection)
	APIMiddleware := middleware.NewAPIMiddleware()

	middleware.Set(r)
	api.Set(r, usecaseCollection, APIMiddleware)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("$PORT must be set")
	}

	HOST := os.Getenv("HOST")
	appPort := fmt.Sprintf("%v:%v", HOST, PORT)

	fmt.Println("Starting the application at:", appPort)
	log.Fatal(http.ListenAndServe(appPort, r))
}
