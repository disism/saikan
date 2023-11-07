package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Dir holds the schema definition for the Dir entity.
type Dir struct {
	ent.Schema
}

// Fields of the Dir.
func (Dir) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name"),
	}
}

// Edges of the Dir.
func (Dir) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("dirs").
			Unique(),
		edge.To("saves", Saved.Type),

		edge.To("subdir", Dir.Type),

		edge.From("pdir", Dir.Type).
			Ref("subdir"),
	}
}

func (Dir) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}
