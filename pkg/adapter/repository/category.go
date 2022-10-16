package repository

import (
	"context"
	"webreader/ent"
	"webreader/ent/category"
	"webreader/pkg/entity/model"
	ur "webreader/pkg/usecase/repository"
)

type categoryRepository struct {
	client *ent.Client
}

// NewCategoryRepository generates  user repository
func NewCategoryRepository(client *ent.Client) ur.Category {
	return &categoryRepository{client: client}
}

func (r *categoryRepository) Get(ctx context.Context, id *model.ID) (*model.Category, error) {
	q := r.client.Category.Query()
	if id != nil {
		q.Where(category.IDEQ(*id))
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

func (r *categoryRepository) List(ctx context.Context) ([]*model.Category, error) {
	ts, err := r.client.
		Category.
		Query().
		All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return ts, nil
}

func (r *categoryRepository) Create(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	u, err := r.client.
		Category.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *categoryRepository) Update(ctx context.Context, input model.UpdateCategoryInput) (*model.Category, error) {
	u, err := r.client.
		Category.UpdateOneID(input.ID).
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
