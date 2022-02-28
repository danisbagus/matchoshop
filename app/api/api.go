package app

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
	handlerV1 "github.com/danisbagus/matchoshop/internal/handler/v1"
	"github.com/danisbagus/matchoshop/internal/repo"

	_ "github.com/lib/pq"
)

func StartApp() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Failed loading .env file")
	}

	client := GetClient()
	defer client.Close()

	router := mux.NewRouter()

	// wiring
	userRepo := repo.NewUserRepo(client)
	productRepo := repo.NewProductRepo(client)
	productCategoryRepo := repo.NewProductCategoryRepo(client)
	productProductCategoryRepo := repo.NewProductProductCategoryRepo(client)
	refreshTokenStoreRepo := repo.NewRefreshTokenStoreRepo(client)

	userService := service.NewUserService(userRepo, refreshTokenStoreRepo)
	productService := service.NewProductService(productRepo, productCategoryRepo, productProductCategoryRepo)
	productCategoryService := service.NewProductCategoryService(productCategoryRepo)

	userHandlerV1 := handlerV1.UserHandler{Service: userService}
	productHandlerV1 := handlerV1.ProductHandler{Service: productService}
	productCategoryHandlerV1 := handlerV1.ProductCategoryHandler{Service: productCategoryService}

	authRouter := router.PathPrefix("/auth").Subrouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	// v1 route
	authRouter.HandleFunc("/v1/login", userHandlerV1.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/v1/refresh", userHandlerV1.Refresh).Methods(http.MethodPost)
	authRouter.HandleFunc("/v1/register/customer", userHandlerV1.RegisterCustomer).Methods(http.MethodPost)

	apiRouter.HandleFunc("/v1/user/profile", userHandlerV1.UpdateUser).Methods(http.MethodPatch)

	apiRouter.HandleFunc("/v1/product", productHandlerV1.CrateProduct).Methods(http.MethodPost)
	apiRouter.HandleFunc("/v1/product", productHandlerV1.GetProductList).Methods(http.MethodGet)
	apiRouter.HandleFunc("/v1/product/{product_id}", productHandlerV1.GetProductDetail).Methods(http.MethodGet)
	apiRouter.HandleFunc("/v1/product/{product_id}", productHandlerV1.UpdateProduct).Methods(http.MethodPut)
	apiRouter.HandleFunc("/v1/product/{product_id}", productHandlerV1.Delete).Methods(http.MethodDelete)

	apiRouter.HandleFunc("/v1/product-category", productCategoryHandlerV1.CrateProductCategory).Methods(http.MethodPost)
	apiRouter.HandleFunc("/v1/product-category", productCategoryHandlerV1.GetProductCategoryList).Methods(http.MethodGet)
	apiRouter.HandleFunc("/v1/product-category/{product_category_id}", productCategoryHandlerV1.GetProductCategoryDetail).Methods(http.MethodGet)
	apiRouter.HandleFunc("/v1/product-category/{product_category_id}", productCategoryHandlerV1.UpdateProductCategory).Methods(http.MethodPut)
	apiRouter.HandleFunc("/v1/product-category/{product_category_id}", productCategoryHandlerV1.Delete).Methods(http.MethodDelete)

	apiRouter.HandleFunc("/hello", SayHello)

	// middleware
	// authMiddleware := middleware.AuthMiddleware{UserRepo: repo.NewUserRepo(client)}
	// apiRouter.Use(authMiddleware.AuthorizationHandler())

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("$PORT must be set")
	}

	HOST := os.Getenv("HOST")

	appPort := fmt.Sprintf("%v:%v", HOST, PORT)
	fmt.Println("Starting the application at:", appPort)
	log.Fatal(http.ListenAndServe(appPort, router))
}

func GetClient() *sqlx.DB {
	dbURL := os.Getenv("DATABASE_URL")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connection := fmt.Sprintf("%s?sslmode=%s", dbURL, dbSSLMode)

	log.Printf("DB url connections: %s", connection)

	client, err := sqlx.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
}
