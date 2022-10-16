package repository

import (
	"context"
	"webreader/pkg/entity/model"
)

// Tag is interface of repository
type Tag interface {
	Get(ctx context.Context, id *model.ID) (*model.Tag, error)
	List(ctx context.Context) ([]*model.Tag, error)
	Create(ctx context.Context, input model.CreateTagInput) (*model.Tag, error)
	Update(ctx context.Context, input model.UpdateTagInput) (*model.Tag, error)
}
