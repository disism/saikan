package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Album holds the schema definition for the Album entity.
type Album struct {
	ent.Schema
}

// Fields of the Album.
func (Album) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("title"),
		field.
			Uint32("year").
			Optional(),
		field.
			String("description").
			Optional(),
	}
}

// Edges of the Album.
func (Album) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("image", Image.Type).
			Ref("albums").
			Required().
			Unique(),
		edge.
			To("musics", Music.Type),
		edge.
			From("users", User.Type).
			Ref("albums"),
		edge.
			From("artists", Artist.Type).
			Ref("albums"),
	}
}

func (Album) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}
