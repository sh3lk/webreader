package controller

import (
	"context"
	"webreader/pkg/entity/model"
	"webreader/pkg/usecase/usecase"
)

// Tag is an interface of controller
type Tag interface {
	Get(ctx context.Context, id *model.ID) (*model.Tag, error)
	List(ctx context.Context) ([]*model.Tag, error)
	Create(ctx context.Context, input model.CreateTagInput) (*model.Tag, error)
	Update(ctx context.Context, input model.UpdateTagInput) (*model.Tag, error)
}

type tag struct {
	tagUsecase usecase.Tag
}

// NewTagController generates test user controller
func NewTagController(uc usecase.Tag) Tag {
	return &tag{
		tagUsecase: uc,
	}
}

func (c *tag) Get(ctx context.Context, id *model.ID) (*model.Tag, error) {
	return c.tagUsecase.Get(ctx, id)
}

func (c *tag) List(ctx context.Context) ([]*model.Tag, error) {
	return c.tagUsecase.List(ctx)
}

func (c *tag) Create(ctx context.Context, input model.CreateTagInput) (*model.Tag, error) {
	return c.tagUsecase.Create(ctx, input)
}

func (c *tag) Update(ctx context.Context, input model.UpdateTagInput) (*model.Tag, error) {
	return c.tagUsecase.Update(ctx, input)
}
