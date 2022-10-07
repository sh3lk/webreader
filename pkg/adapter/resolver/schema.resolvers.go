package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"webreader/ent"
	"webreader/ent/schema/ulid"
	"webreader/graph/generated"
)

func (r *queryResolver) Node(ctx context.Context, id ulid.ID) (ent.Noder, error) {
	n, err := r.client.Noder(ctx, id, ent.WithNodeType(ent.IDToType))
	if err != nil {
		return nil, err
	}
	return n, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
