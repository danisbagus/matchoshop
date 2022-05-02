package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"github.com/danisbagus/matchoshop/app/api/middleware"
	"github.com/danisbagus/matchoshop/internal/core/service"
	handlerV1 "github.com/danisbagus/matchoshop/internal/handler/v1"
	"github.com/danisbagus/matchoshop/internal/repo"
	"github.com/danisbagus/matchoshop/utils/constants"

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
	orderRepo := repo.NewOrderRepo(client)
	orderProductRepo := repo.NewOrderProductRepo(client)
	paymentResultRepo := repo.NewPaymentResult(client)

	userService := service.NewUserService(userRepo, refreshTokenStoreRepo)
	productService := service.NewProductService(productRepo, productCategoryRepo, productProductCategoryRepo)
	productCategoryService := service.NewProductCategoryService(productCategoryRepo)
	orderService := service.NewOrderService(orderRepo, orderProductRepo, paymentResultRepo, userRepo)
	uploadService := service.NewUploadService()

	userHandlerV1 := handlerV1.UserHandler{Service: userService}
	productHandlerV1 := handlerV1.ProductHandler{Service: productService}
	productCategoryHandlerV1 := handlerV1.ProductCategoryHandler{Service: productCategoryService}
	orderHandlerV1 := handlerV1.OrderHandler{Service: orderService}
	configHandlerV1 := handlerV1.ConfigHandler{}
	uploadHandlerV1 := handlerV1.UploadHandler{Service: uploadService}

	// auth v1 routes
	authV1Route := router.PathPrefix("/api/v1/auth").Subrouter()
	authV1Route.HandleFunc("/login", userHandlerV1.Login).Methods(http.MethodPost)
	authV1Route.HandleFunc("/refresh", userHandlerV1.Refresh).Methods(http.MethodPost)
	authV1Route.HandleFunc("/register/customer", userHandlerV1.RegisterCustomer).Methods(http.MethodPost)

	// user v1 routes
	userV1Route := router.PathPrefix("/api/v1/user").Subrouter()
	userV1Route.Use(middleware.AuthorizationHandler())
	userV1Route.HandleFunc("", userHandlerV1.GetUserDetail).Methods(http.MethodGet)
	userV1Route.HandleFunc("/profile", userHandlerV1.UpdateUser).Methods(http.MethodPatch)

	// product v1 routes
	productV1Route := router.PathPrefix("/api/v1/product").Subrouter()
	productV1Route.HandleFunc("", productHandlerV1.GetProductList).Methods(http.MethodGet)
	productV1Route.HandleFunc("/{product_id}", productHandlerV1.GetProductDetail).Methods(http.MethodGet)

	// product admin v1 routes
	productAdminV1Route := router.PathPrefix("/api/v1/admin/product").Subrouter()
	productAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	productAdminV1Route.HandleFunc("", productHandlerV1.CreateProduct).Methods(http.MethodPost)
	productAdminV1Route.HandleFunc("/{product_id}", productHandlerV1.UpdateProduct).Methods(http.MethodPut)
	productAdminV1Route.HandleFunc("/{product_id}", productHandlerV1.Delete).Methods(http.MethodDelete)
	productAdminV1Route.HandleFunc("/{product_id}", productHandlerV1.GetProductDetail).Methods(http.MethodGet)

	// product category v1 routes
	productCategoryV1Route := router.PathPrefix("/api/v1/product-category").Subrouter()
	productCategoryV1Route.HandleFunc("", productCategoryHandlerV1.GetProductCategoryList).Methods(http.MethodGet)
	productCategoryV1Route.HandleFunc("/{product_category_id}", productCategoryHandlerV1.GetProductCategoryDetail).Methods(http.MethodGet)

	// product category admin v1 routes
	productCategoryAdminV1Route := router.PathPrefix("/api/v1/admin/product-category").Subrouter()
	productAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	productCategoryAdminV1Route.HandleFunc("", productCategoryHandlerV1.CreateProductCategory).Methods(http.MethodPost)
	productCategoryAdminV1Route.HandleFunc("/{product_category_id}", productCategoryHandlerV1.UpdateProductCategory).Methods(http.MethodPut)
	productCategoryAdminV1Route.HandleFunc("/{product_category_id}", productCategoryHandlerV1.Delete).Methods(http.MethodDelete)

	// order v1 routes
	orderV1Route := router.PathPrefix("/api/v1/order").Subrouter()
	orderV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.CustomerPermission))
	orderV1Route.HandleFunc("", orderHandlerV1.Create).Methods(http.MethodPost)
	orderV1Route.HandleFunc("", orderHandlerV1.GetList).Methods(http.MethodGet)
	orderV1Route.HandleFunc("/{order_id}", orderHandlerV1.GetDetail).Methods(http.MethodGet)
	orderV1Route.HandleFunc("/{order_id}/pay", orderHandlerV1.UpdatePaid).Methods(http.MethodPut)

	// product admin v1 routes
	userAdminV1Route := router.PathPrefix("/api/v1/admin/user").Subrouter()
	userAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	userAdminV1Route.HandleFunc("", userHandlerV1.GetUserList).Methods(http.MethodGet)
	userAdminV1Route.HandleFunc("/{user_id}", userHandlerV1.GetUserDetailAdmin).Methods(http.MethodGet)
	userAdminV1Route.HandleFunc("/{user_id}", userHandlerV1.DeleteUser).Methods(http.MethodDelete)
	userAdminV1Route.HandleFunc("/{user_id}", userHandlerV1.UpdateUserAdmin).Methods(http.MethodPatch)

	// config routes
	configRoute := router.PathPrefix("/api/v1/config").Subrouter()
	configRoute.HandleFunc("/paypal", configHandlerV1.GetPaypalConfig).Methods(http.MethodGet)

	// upload
	uploadRoute := router.PathPrefix("/api/v1/upload").Subrouter()
	uploadRoute.HandleFunc("/image", uploadHandlerV1.UploadImage).Methods(http.MethodPost)

	orderAdminV1Route := router.PathPrefix("/api/v1/admin/order").Subrouter()
	orderAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	orderAdminV1Route.HandleFunc("", orderHandlerV1.GetListAdmin).Methods(http.MethodGet)
	orderAdminV1Route.HandleFunc("/{order_id}", orderHandlerV1.GetDetail).Methods(http.MethodGet)
	orderAdminV1Route.HandleFunc("/{order_id}/deliver", orderHandlerV1.UpdateDelivered).Methods(http.MethodPut)

	router.HandleFunc("/health-check", healthCheck)

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

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("App Up"))
}
