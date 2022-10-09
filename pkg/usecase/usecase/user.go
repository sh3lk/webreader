package usecase

import (
	"context"
	"webreader/pkg/entity/model"
	"webreader/pkg/usecase/repository"
)

type user struct {
	userRepository repository.User
}

// User of usecase.
type User interface {
	Get(ctx context.Context, id *model.ID) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	CreateWithTodo(ctx context.Context, input model.CreateUserInput) (*model.User, error)
	Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error)
}

// NewUserUsecase returns user usecase.
func NewUserUsecase(r repository.User) User {
	return &user{userRepository: r}
}

func (u *user) Get(ctx context.Context, id *model.ID) (*model.User, error) {
	return u.userRepository.Get(ctx, id)
}

func (u *user) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.userRepository.GetByEmail(ctx, email)
}

func (u *user) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.UserWhereInput) (*model.UserConnection, error) {
	return u.userRepository.List(ctx, after, first, before, last, where)
}

func (u *user) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return u.userRepository.Create(ctx, input)
}

func (u *user) CreateWithTodo(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return u.userRepository.CreateWithTodo(ctx, input)
}

func (u *user) Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return u.userRepository.Update(ctx, input)
}
