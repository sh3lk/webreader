package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"webreader/ent"
	"webreader/ent/schema/ulid"
)

// CreateRanobe is the resolver for the createRanobe field.
func (r *mutationResolver) CreateRanobe(ctx context.Context, input ent.CreateRanobeInput) (*ent.Ranobe, error) {
	return r.client.Ranobe.Create().SetInput(input).Save(ctx)
}

// UpdateRanobe is the resolver for the updateRanobe field.
func (r *mutationResolver) UpdateRanobe(ctx context.Context, input ent.UpdateRanobeInput) (*ent.Ranobe, error) {
	panic(fmt.Errorf("not implemented: UpdateRanobe - updateRanobe"))
}

// Ranobe is the resolver for the ranobe field.
func (r *queryResolver) Ranobe(ctx context.Context, id *ulid.ID) (*ent.Ranobe, error) {
	panic(fmt.Errorf("not implemented: Ranobe - ranobe"))
}

// Ranobes is the resolver for the ranobes field.
func (r *queryResolver) Ranobes(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.RanobeWhereInput) (*ent.UserConnection, error) {
	panic(fmt.Errorf("not implemented: Ranobes - ranobes"))
}
