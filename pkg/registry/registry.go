package registry

import (
	"webreader/ent"
	"webreader/pkg/adapter/controller"
)

type registry struct {
	client *ent.Client
}

// Registry is an interface of registry
type Registry interface {
	NewController() controller.Controller
}

// New registers entire controller with dependencies
func New(client *ent.Client) Registry {
	return &registry{
		client: client,
	}
}

// NewController generates controllers
func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		User:     r.NewUserController(),
		Todo:     r.NewTodoController(),
		Category: r.NewCategoryController(),
		Tag:      r.NewTagController(),
	}
}
