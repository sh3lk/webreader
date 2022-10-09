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

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
	t, err := r.controller.Todo.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input ent.UpdateTodoInput) (*ent.Todo, error) {
	t, err := r.controller.Todo.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

// Todo is the resolver for the Todo field.
func (r *queryResolver) Todo(ctx context.Context, id *ulid.ID) (*ent.Todo, error) {
	t, err := r.controller.Todo.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

// Todos is the resolver for the Todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	ts, err := r.controller.Todo.List(ctx)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return ts, nil
}

// CreatedAt is the resolver for the createdAt field.
func (r *todoResolver) CreatedAt(ctx context.Context, obj *ent.Todo) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *todoResolver) UpdatedAt(ctx context.Context, obj *ent.Todo) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
