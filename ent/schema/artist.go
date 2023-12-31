package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Artist holds the schema definition for the Artist entity.
type Artist struct {
	ent.Schema
}

// Fields of the Artist.
func (Artist) Fields() []ent.Field {
	return []ent.Field{
		field.
			Uint64("id"),
		field.
			String("name").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Artist.
func (Artist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("musics", Music.Type),
		edge.
			To("albums", Album.Type),
	}
}
