### Mockery Update
### example : make mockery-generate-repo repoInterface=IUserRepository mockFile=user_mock
mockery-generate-repo:
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name $(repoInterface) --filename $(mockFile).go
	@echo "$(repoInterface) mocks generated"

mockery-generate-repo-all:
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name IHealthcheckRepository --filename healthcheck_mock.go
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name IOrderRepository --filename order_mock.go
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name IOrderProductRepository --filename order_product_mock.go
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name IPaymentResultRepository --filename payment_result_mock.go
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name IProductCategoryRepository --filename product_category_mock.go
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name IProductProductCategoryRepository --filename product_product_category_mock.go
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name IProductRepository --filename product_mock.go
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name IRefreshTokenStoreRepository --filename refresh_token_store_mock.go
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name IReviewRepository --filename review_mock.go
	mockery --dir ./internal/repository --output ./internal/repository/mocks --name IUserRepository --filename user_mock.go
	@echo "mocks generated"