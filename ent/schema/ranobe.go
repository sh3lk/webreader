package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
	"webreader/ent/mixin"
	"webreader/pkg/const/globalid"
)

// RanobeMixin defines Fields
type RanobeMixin struct {
	entMixin.Schema
}

// Ranobe holds the schema definition for the Ranobe entity.
type Ranobe struct {
	ent.Schema
}

// Fields of the Ranobe.
func (Ranobe) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("engName"),
		field.String("rusName"),
		field.Text("summary"),
		field.Int("releaseDate"),
	}
}

// Edges of the Ranobe.
func (Ranobe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("categories", Category.Type),
		edge.To("tags", Tag.Type),
	}
}

// Mixin of the Ranobe.
func (Ranobe) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Ranobe.Prefix),
		RanobeMixin{},
		mixin.NewDatetime(),
	}
}
