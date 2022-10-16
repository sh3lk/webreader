package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *categoryResolver) Categories(ctx context.Context, obj *ent.Category) ([]*ent.Category, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}
