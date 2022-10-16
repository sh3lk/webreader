package usecase

import (
	"context"
	"webreader/pkg/entity/model"
	"webreader/pkg/usecase/repository"
)

type tag struct {
	tagRepository repository.Tag
}

// Tag is an interface of test user
type Tag interface {
	Get(ctx context.Context, id *model.ID) (*model.Tag, error)
	List(ctx context.Context) ([]*model.Tag, error)
	Create(ctx context.Context, input model.CreateTagInput) (*model.Tag, error)
	Update(ctx context.Context, input model.UpdateTagInput) (*model.Tag, error)
}

// NewTagUsecase generates test user repository
func NewTagUsecase(r repository.Tag) Tag {
	return &tag{tagRepository: r}
}

func (t *tag) Get(ctx context.Context, id *model.ID) (*model.Tag, error) {
	return t.tagRepository.Get(ctx, id)
}

func (t *tag) List(ctx context.Context) ([]*model.Tag, error) {
	return t.tagRepository.List(ctx)
}

func (t *tag) Create(ctx context.Context, input model.CreateTagInput) (*model.Tag, error) {
	return t.tagRepository.Create(ctx, input)
}

func (t *tag) Update(ctx context.Context, input model.UpdateTagInput) (*model.Tag, error) {
	return t.tagRepository.Update(ctx, input)
}
