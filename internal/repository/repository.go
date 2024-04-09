package repository

import (
	"github.com/jmoiron/sqlx"
)

type RepositoryCollection struct {
	HealthCheckRepository            IHealthcheckRepository
	OrderRepository                  IOrderRepository
	OrderProductRepository           IOrderProductRepository
	PaymentResultRepository          IPaymentResultRepository
	ProductReposotory                IProductRepository
	ProductCategoryRepository        IProductCategoryRepository
	ProductProductCategoryRepository IProductProductCategoryRepository
	RefreshTokenStoreRepository      IRefreshTokenStoreRepository
	ReviewRepository                 IReviewRepository
	UserRepository                   IUserRepository
}

func NewRepoCollection(db *sqlx.DB) RepositoryCollection {
	return RepositoryCollection{
		HealthCheckRepository:            NewHealthCheckRepository(db),
		OrderRepository:                  NewOrderRepository(db),
		OrderProductRepository:           NewOrderProductRepository(db),
		PaymentResultRepository:          NewPaymentResultRepository(db),
		ProductReposotory:                NewProductReposotory(db),
		ProductCategoryRepository:        NewProductCategoryRepository(db),
		ProductProductCategoryRepository: NewProductProductCategoryRepository(db),
		RefreshTokenStoreRepository:      NewRefreshTokenStoreRepository(db),
		ReviewRepository:                 NewReviewRepository(db),
		UserRepository:                   NewUserRepository(db),
	}
}
