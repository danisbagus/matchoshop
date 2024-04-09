package mocks

type RepoCollectionMocks struct {
	HealthCheckRepository            *IHealthcheckRepository
	OrderRepository                  *IOrderRepository
	OrderProductRepository           *IOrderProductRepository
	PaymentResultRepository          *IPaymentResultRepository
	ProductReposotory                *IProductRepository
	ProductCategoryRepository        *IProductCategoryRepository
	ProductProductCategoryRepository *IProductProductCategoryRepository
	RefreshTokenStoreRepository      *IRefreshTokenStoreRepository
	ReviewRepository                 *IReviewRepository
	UserRepository                   *IUserRepository
}
