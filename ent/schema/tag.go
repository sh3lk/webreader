package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
	"webreader/ent/mixin"
	"webreader/pkg/const/globalid"
)

// TagMixin defines Fields
type TagMixin struct {
	entMixin.Schema
}

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty(),
		field.Text("description"),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ranobes", Ranobe.Type).
			Ref("tags"),
	}
}

// Mixin of the Todo.
func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Tag.Prefix),
		TagMixin{},
		mixin.NewDatetime(),
	}
}
