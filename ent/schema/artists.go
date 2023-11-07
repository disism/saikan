package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Artists holds the schema definition for the Artists entity.
type Artists struct {
	ent.Schema
}

// Fields of the Artists.
func (Artists) Fields() []ent.Field {
	return []ent.Field{
		field.
			Uint64("id"),
		field.
			String("name").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Artists.
func (Artists) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("musics", Musics.Type),
		edge.
			To("albums", Albums.Type),
	}
}
