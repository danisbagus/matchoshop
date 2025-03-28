package api

import (
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"

	"github.com/joho/godotenv"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/app/api/middleware"
	"github.com/danisbagus/matchoshop/internal/core/service"
	handlerV1 "github.com/danisbagus/matchoshop/internal/handler/v1"
	"github.com/danisbagus/matchoshop/internal/repo"
	"github.com/danisbagus/matchoshop/utils/constants"
	"github.com/danisbagus/matchoshop/utils/modules"

	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func StartApp() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed loading .env file")
	}

	e := echo.New()

	e.Use(
		echoMid.BodyDumpWithConfig(echoMid.BodyDumpConfig{
			Handler: logHandler,
		}),
		echoMid.Recover(),
		echoMid.CORSWithConfig(echoMid.CORSConfig{
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		}),
	)

	client := modules.GetPostgresClient()

	// wiring
	userRepo := repo.NewUserRepo(client)
	productRepo := repo.NewProductRepo(client)
	productCategoryRepo := repo.NewProductCategoryRepo(client)
	productProductCategoryRepo := repo.NewProductProductCategoryRepo(client)
	refreshTokenStoreRepo := repo.NewRefreshTokenStoreRepo(client)
	orderRepo := repo.NewOrderRepo(client)
	orderProductRepo := repo.NewOrderProductRepo(client)
	paymentResultRepo := repo.NewPaymentResult(client)
	reviewRepo := repo.NewReviewRepo(client)
	healthCheckRepo := repo.NewHealthCheck(client)

	userService := service.NewUserService(userRepo, refreshTokenStoreRepo)
	productService := service.NewProductService(productRepo, productCategoryRepo, productProductCategoryRepo, reviewRepo)
	productCategoryService := service.NewProductCategoryService(productCategoryRepo)
	orderService := service.NewOrderService(orderRepo, orderProductRepo, paymentResultRepo, productRepo)
	uploadService := service.NewUploadService()
	reviewService := service.NewReviewService(reviewRepo)
	healthCheckService := service.NewHealthCheckService(healthCheckRepo)

	userHandlerV1 := handlerV1.NewUserhandler(userService)
	productHandlerV1 := handlerV1.NewProductHandler(productService)
	productCategoryHandlerV1 := handlerV1.NewProductCategoryHandler(productCategoryService)
	orderHandlerV1 := handlerV1.NewOrderHandler(orderService)
	reviewHandlerV1 := handlerV1.NewReviewHandler(reviewService)
	uploadHandlerV1 := handlerV1.NewUploadHandler(uploadService)
	healthCheckHandlerV1 := handlerV1.NewHealthCheckHandlerHandler(healthCheckService)

	// auth v1 routes
	authV1Route := e.Group("/api/v1/auth")
	authV1Route.POST("/login", userHandlerV1.Login)
	authV1Route.POST("/refresh", userHandlerV1.Refresh)
	authV1Route.POST("/register/customer", userHandlerV1.RegisterCustomer)

	// user v1 routes
	userV1Route := e.Group("/api/v1/user")
	userV1Route.Use(middleware.AuthorizationHandler())
	userV1Route.GET("", userHandlerV1.GetUserDetail)
	userV1Route.PATCH("/profile", userHandlerV1.UpdateUser)

	// user admin v1 routes
	userAdminV1Route := e.Group("/api/v1/admin/user")
	userAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	userAdminV1Route.GET("", userHandlerV1.GetUserList)
	userAdminV1Route.GET("/:user_id", userHandlerV1.GetUserDetailAdmin)
	userAdminV1Route.DELETE("/:user_id", userHandlerV1.DeleteUser)
	userAdminV1Route.PATCH("/:user_id", userHandlerV1.UpdateUserAdmin)

	// product v1 routes
	productV1Route := e.Group("/api/v1/product")
	productV1Route.GET("", productHandlerV1.GetProductListPaginate)
	productV1Route.GET("/top", productHandlerV1.GetTopProduct)
	productV1Route.GET("/:product_id", productHandlerV1.GetProductDetail)

	// product admin v1 routes
	productAdminV1Route := e.Group("/api/v1/admin/product")
	productAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	productAdminV1Route.POST("", productHandlerV1.CreateProduct)
	productAdminV1Route.GET("", productHandlerV1.GetProductListPaginate)
	productAdminV1Route.PUT("/:product_id", productHandlerV1.UpdateProduct)
	productAdminV1Route.DELETE("/:product_id", productHandlerV1.Delete)
	productAdminV1Route.GET("/:product_id", productHandlerV1.GetProductDetail)

	// product category v1 routes
	productCategoryV1Route := e.Group("/api/v1/product-category")
	productCategoryV1Route.GET("", productCategoryHandlerV1.GetProductCategoryList)
	productCategoryV1Route.GET("/:product_category_id", productCategoryHandlerV1.GetProductCategoryDetail)

	// product category admin v1 routes
	productCategoryAdminV1Route := e.Group("/api/v1/admin/product-category")
	productCategoryAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	productCategoryAdminV1Route.POST("", productCategoryHandlerV1.CreateProductCategory)
	productCategoryAdminV1Route.PUT("/:product_category_id", productCategoryHandlerV1.UpdateProductCategory)
	productCategoryAdminV1Route.DELETE("/:product_category_id", productCategoryHandlerV1.Delete)

	// order v1 routes
	orderV1Route := e.Group("/api/v1/order")
	orderV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.CustomerPermission))
	orderV1Route.POST("", orderHandlerV1.Create)
	orderV1Route.GET("", orderHandlerV1.GetList)
	orderV1Route.GET("/:order_id", orderHandlerV1.GetDetail)
	orderV1Route.PUT("/:order_id/pay", orderHandlerV1.UpdatePaid)

	// order admin v1 routes
	orderAdminV1Route := e.Group("/api/v1/admin/order")
	orderAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	orderV1Route.GET("", orderHandlerV1.GetListAdmin)
	orderV1Route.GET("/:order_id", orderHandlerV1.GetDetail)
	orderV1Route.PUT("/:order_id/deliver", orderHandlerV1.UpdateDelivered)

	// review v1 routes
	reviewV1Route := e.Group("/api/v1/review")
	reviewV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.CustomerPermission))
	reviewV1Route.POST("", reviewHandlerV1.Create)
	reviewV1Route.PUT("", reviewHandlerV1.Update)
	reviewV1Route.GET("/:product_id", reviewHandlerV1.GetDetail)

	// upload v1 routes
	uploadRoute := e.Group("/api/v1/upload")
	uploadRoute.POST("/image", uploadHandlerV1.UploadImage)

	// health check v1 router
	healthCheckRouter := e.Group("/api/v1/health-check")
	healthCheckRouter.GET("", healthCheckHandlerV1.Get)
	healthCheckRouter.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	appPort := os.Getenv("PORT") // todo: move to config
	e.Logger.Fatal(e.Start(":" + appPort))
}

func logHandler(c echo.Context, req, res []byte) { // todo: move to utils
	logger.Info("incoming request",
		zap.String("request_method", c.Request().Method),
		zap.String("request_uri", c.Request().RequestURI),
		zap.String("request_data", string(req)),
		zap.String("response_data", string(res)),
	)
}
