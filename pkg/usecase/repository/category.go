package repository

import (
	"context"
	"webreader/pkg/entity/model"
)

// Category is interface of repository
type Category interface {
	Get(ctx context.Context, id *model.ID) (*model.Category, error)
	List(ctx context.Context) ([]*model.Category, error)
	Create(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error)
	Update(ctx context.Context, input model.UpdateCategoryInput) (*model.Category, error)
}
