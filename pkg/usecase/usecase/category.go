package usecase

import (
	"context"
	"webreader/pkg/entity/model"
	"webreader/pkg/usecase/repository"
)

type category struct {
	categoryRepository repository.Category
}

// Category is an interface of test user
type Category interface {
	Get(ctx context.Context, id *model.ID) (*model.Category, error)
	List(ctx context.Context) ([]*model.Category, error)
	Create(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error)
	Update(ctx context.Context, input model.UpdateCategoryInput) (*model.Category, error)
}

// NewCategoryUsecase generates test user repository
func NewCategoryUsecase(r repository.Category) Category {
	return &category{categoryRepository: r}
}

func (c *category) Get(ctx context.Context, id *model.ID) (*model.Category, error) {
	return c.categoryRepository.Get(ctx, id)
}

func (c *category) List(ctx context.Context) ([]*model.Category, error) {
	return c.categoryRepository.List(ctx)
}

func (c *category) Create(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	return c.categoryRepository.Create(ctx, input)
}

func (c *category) Update(ctx context.Context, input model.UpdateCategoryInput) (*model.Category, error) {
	return c.categoryRepository.Update(ctx, input)
}
