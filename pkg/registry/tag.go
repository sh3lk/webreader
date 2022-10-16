package registry

import (
	"webreader/pkg/adapter/controller"
	"webreader/pkg/adapter/repository"
	"webreader/pkg/usecase/usecase"
)

// NewTagController conforms to interface
func (r *registry) NewTagController() controller.Tag {
	repo := repository.NewTagRepository(r.client)
	u := usecase.NewTagUsecase(repo)

	return controller.NewTagController(u)
}
