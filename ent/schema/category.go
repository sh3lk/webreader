package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
	"webreader/ent/mixin"
	"webreader/pkg/const/globalid"
)

// CategoryMixin defines Fields
type CategoryMixin struct {
	entMixin.Schema
}

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.Text("description"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ranobes", Ranobe.Type).
			Ref("categories"),
	}
}

// Mixin of the Todo.
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Category.Prefix),
		CategoryMixin{},
		mixin.NewDatetime(),
	}
}
