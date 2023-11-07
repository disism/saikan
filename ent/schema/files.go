package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Files holds the schema definition for the Files entity.
type Files struct {
	ent.Schema
}

// Fields of the Files.
func (Files) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("hash").
			Unique(),
		field.
			String("name").
			NotEmpty(),
		field.
			Uint64("size"),
	}
}

// Edges of the Files.
func (Files) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("musics", Musics.Type),
		edge.
			To("images", Images.Type),
	}
}

func (Files) Indexes() []ent.Index {
	return []ent.Index{
		index.
			Fields("hash").
			Unique(),
	}
}

func (Files) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
