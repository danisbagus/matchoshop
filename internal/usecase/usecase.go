package usecase

import "github.com/danisbagus/matchoshop/internal/repository"

type UsecaseCollection struct {
	HealthcheckUsecase     IHealthcheckUsecase
	OrderUsecase           IOrderUsecase
	ProductCategoryUsecase IProductCategoryUsecase
	ProductUsecase         IProductUsecase
	ReviewUsecase          IReviewUsecase
	UploadUsecase          IUploadUsecase
	UserUsecase            IUserUsecase
}

func NewUsecaseCollection(repositoryCollection repository.RepositoryCollection) UsecaseCollection {
	return UsecaseCollection{
		HealthcheckUsecase:     NewHealthcheckUsecase(repositoryCollection),
		OrderUsecase:           NewOrderUsecase(repositoryCollection),
		ProductCategoryUsecase: NewProductCategoryUsecase(repositoryCollection),
		ProductUsecase:         NewProductUsecase(repositoryCollection),
		ReviewUsecase:          NewReviewUsecase(repositoryCollection),
		UploadUsecase:          NewUploadUsecase(),
		UserUsecase:            NewUserUsecase(repositoryCollection),
	}
}
