package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"github.com/danisbagus/matchoshop/internal/core/service"
	"github.com/danisbagus/matchoshop/internal/handler"
	"github.com/danisbagus/matchoshop/internal/repo"

	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := GetClient()
	defer client.Close()

	router := mux.NewRouter()

	userRepo := repo.NewUserRepo(client)

	userService := service.NewUserService(userRepo)

	userHandler := handler.AuthHandler{Service: userService}

	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/register/merchant", userHandler.RegisterMerchant).Methods(http.MethodPost)

	appPort := fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	fmt.Println("Starting the application at:", appPort)
	log.Fatal(http.ListenAndServe(appPort, router))
}

func GetClient() *sqlx.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUser, dbPassword, dbHost, dbPort, dbName)
	// connection := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
	// 	dbHost, dbPort, dbUser, dbPassword, dbName)

	// goose -dir db/migration postgres "postgres://postgres:mypass@localhost:8010/matchoshop?sslmode=disable" up

	connection := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	fmt.Println(connection)

	client, err := sqlx.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
