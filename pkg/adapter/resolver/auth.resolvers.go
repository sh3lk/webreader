package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"webreader/ent"
	"webreader/ent/schema/ulid"
	"webreader/pkg/adapter/handler"
	"webreader/pkg/entity/model"
	"webreader/pkg/util/auth"

	"github.com/golang-jwt/jwt"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input ent.CreateUserInput) (*model.Credential, error) {
	u, err := r.controller.User.GetByEmail(ctx, input.Email)

	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	if !auth.CheckPasswordHash(input.Password, u.Password) {
		return nil, handler.HandleError(
			ctx,
			errors.New("incorrect credentials"),
		)
	}

	t, err := auth.GenerateToken(u.ID)

	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	credential := &model.Credential{
		Target: u,
		Token:  t,
	}

	return credential, nil
}

// Registration is the resolver for the Registration field.
func (r *mutationResolver) Registration(ctx context.Context, input ent.CreateUserInput) (*model.Credential, error) {
	hash, err := auth.HashPassword(input.Password)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	input.Password = hash
	u, err := r.controller.User.Create(ctx, input)

	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	t, err := auth.GenerateToken(u.ID)

	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	credential := &model.Credential{
		Target: u,
		Token:  t,
	}

	return credential, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	auth := ctx.Value("auth")

	if auth == nil {
		return nil, handler.HandleError(
			ctx,
			errors.New("incorrect token"),
		)
	}
	uid := auth.(ulid.ID)
	u, err := r.controller.User.Get(ctx, &uid)

	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	return u, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
type jwtCustomClaims struct {
	Id ulid.ID `json:"id"`
	jwt.StandardClaims
}
