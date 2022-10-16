package repository

import (
	"context"
	"webreader/ent"
	"webreader/ent/tag"
	"webreader/pkg/entity/model"
	ur "webreader/pkg/usecase/repository"
)

type tagRepository struct {
	client *ent.Client
}

// NewTagRepository generates  user repository
func NewTagRepository(client *ent.Client) ur.Tag {
	return &tagRepository{client: client}
}

func (r *tagRepository) Get(ctx context.Context, id *model.ID) (*model.Tag, error) {
	q := r.client.Tag.Query()
	if id != nil {
		q.Where(tag.IDEQ(*id))
	}

	u, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, map[string]interface{}{
				"id": id,
			})
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *tagRepository) List(ctx context.Context) ([]*model.Tag, error) {
	ts, err := r.client.
		Tag.
		Query().
		All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return ts, nil
}

func (r *tagRepository) Create(ctx context.Context, input model.CreateTagInput) (*model.Tag, error) {
	u, err := r.client.
		Tag.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *tagRepository) Update(ctx context.Context, input model.UpdateTagInput) (*model.Tag, error) {
	u, err := r.client.
		Tag.UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return u, nil
}
