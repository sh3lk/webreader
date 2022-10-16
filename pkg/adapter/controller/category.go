package controller

import (
	"context"
	"webreader/pkg/entity/model"
	"webreader/pkg/usecase/usecase"
)

// Category is an interface of controller
type Category interface {
	Get(ctx context.Context, id *model.ID) (*model.Category, error)
	List(ctx context.Context) ([]*model.Category, error)
	Create(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error)
	Update(ctx context.Context, input model.UpdateCategoryInput) (*model.Category, error)
}

type category struct {
	categoryUsecase usecase.Category
}

// NewCategoryController generates test user controller
func NewCategoryController(uc usecase.Category) Category {
	return &category{
		categoryUsecase: uc,
	}
}

func (c *category) Get(ctx context.Context, id *model.ID) (*model.Category, error) {
	return c.categoryUsecase.Get(ctx, id)
}

func (c *category) List(ctx context.Context) ([]*model.Category, error) {
	return c.categoryUsecase.List(ctx)
}

func (c *category) Create(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	return c.categoryUsecase.Create(ctx, input)
}

func (c *category) Update(ctx context.Context, input model.UpdateCategoryInput) (*model.Category, error) {
	return c.categoryUsecase.Update(ctx, input)
}
