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

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input ent.CreateTagInput) (*ent.Tag, error) {
	t, err := r.controller.Tag.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

// UpdateTag is the resolver for the updateTag field.
func (r *mutationResolver) UpdateTag(ctx context.Context, input ent.UpdateTagInput) (*ent.Tag, error) {
	t, err := r.controller.Tag.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

// Tag is the resolver for the tag field.
func (r *queryResolver) Tag(ctx context.Context, id *ulid.ID) (*ent.Tag, error) {
	t, err := r.controller.Tag.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context) ([]*ent.Tag, error) {
	ts, err := r.controller.Tag.List(ctx)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return ts, nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *tagResolver) CreatedAt(ctx context.Context, obj *ent.Tag) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *tagResolver) UpdatedAt(ctx context.Context, obj *ent.Tag) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// Tag returns generated.TagResolver implementation.
func (r *Resolver) Tag() generated.TagResolver { return &tagResolver{r} }

type tagResolver struct{ *Resolver }
