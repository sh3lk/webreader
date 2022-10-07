package registry

import (
	"webreader/pkg/adapter/controller"
	"webreader/pkg/adapter/repository"
	"webreader/pkg/usecase/usecase"
)

// NewUserController conforms to interface
func (r *registry) NewUserController() controller.User {
	repo := repository.NewUserRepository(r.client)
	u := usecase.NewUserUsecase(repo)

	return controller.NewUserController(u)
}
