package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"webreader/ent"
	"webreader/ent/schema/ulid"
	"webreader/graph/generated"
	"webreader/pkg/adapter/handler"
	"webreader/pkg/util/datetime"
)

// CreatedAt is the resolver for the createdAt field.
func (r *categoryResolver) CreatedAt(ctx context.Context, obj *ent.Category) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *categoryResolver) UpdatedAt(ctx context.Context, obj *ent.Category) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input ent.CreateCategoryInput) (*ent.Category, error) {
	c, err := r.controller.Category.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return c, nil
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, input ent.UpdateCategoryInput) (*ent.Category, error) {
	c, err := r.controller.Category.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return c, nil
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id *ulid.ID) (*ent.Category, error) {
	c, err := r.controller.Category.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return c, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*ent.Category, error) {
	cs, err := r.controller.Category.List(ctx)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return cs, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

type categoryResolver struct{ *Resolver }
