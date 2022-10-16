// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (c *Category) Ranobes(ctx context.Context) (result []*Ranobe, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedRanobes(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.RanobesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryRanobes().All(ctx)
	}
	return result, err
}

func (r *Ranobe) Categories(ctx context.Context) (result []*Category, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedCategories(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.CategoriesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryCategories().All(ctx)
	}
	return result, err
}

func (t *Todo) User(ctx context.Context) (*User, error) {
	result, err := t.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Todos(ctx context.Context) (result []*Todo, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedTodos(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.TodosOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryTodos().All(ctx)
	}
	return result, err
}
