package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Albums holds the schema definition for the Albums entity.
type Albums struct {
	ent.Schema
}

// Fields of the Albums.
func (Albums) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("title"),
		field.
			String("date").
			MaxLen(20),
		field.
			String("description").
			Optional(),
	}
}

// Edges of the Albums.
func (Albums) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("image", Images.Type).
			Ref("albums").
			Required().
			Unique(),
		edge.
			To("musics", Musics.Type),
		edge.
			From("users", Users.Type).
			Ref("albums"),
		edge.
			From("artists", Artists.Type).
			Ref("albums"),
	}
}

func (Albums) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}
