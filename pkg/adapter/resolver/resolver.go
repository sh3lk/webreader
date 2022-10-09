package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"webreader/ent"
	"webreader/graph/generated"
	"webreader/pkg/adapter/controller"
)

// Resolver is a context struct
type Resolver struct {
	client     *ent.Client
	controller controller.Controller
}

// NewSchema creates NewExecutableSchema
func NewSchema(client *ent.Client, controller controller.Controller) graphql.ExecutableSchema {
	config := generated.Config{
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
		},
	}

	config.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		if ctx.Value("auth") == nil {
			// block calling the next resolver
			return nil, errors.New("Access denied")
		}

		return next(ctx)
	}

	return generated.NewExecutableSchema(config)
}
