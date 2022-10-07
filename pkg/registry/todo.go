package registry

import (
	"webreader/pkg/adapter/controller"
	"webreader/pkg/adapter/repository"
	"webreader/pkg/usecase/usecase"
)

func (r *registry) NewTodoController() controller.Todo {
	repo := repository.NewTodoRepository(r.client)
	u := usecase.NewTodoUsecase(repo)

	return controller.NewTodoController(u)
}
