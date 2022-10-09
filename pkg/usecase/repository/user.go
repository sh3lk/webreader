package repository

import (
	"context"
	"webreader/pkg/entity/model"
)

// User is interface of repository
type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	CreateWithTodo(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error)
}
