package registry

import (
	"webreader/pkg/adapter/controller"
	"webreader/pkg/adapter/repository"
	"webreader/pkg/usecase/usecase"
)

// NewCategoryController conforms to interface
func (r *registry) NewCategoryController() controller.Category {
	repo := repository.NewCategoryRepository(r.client)
	u := usecase.NewCategoryUsecase(repo)

	return controller.NewCategoryController(u)
}
