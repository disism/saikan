package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Musics holds the schema definition for the Musics entity.
type Musics struct {
	ent.Schema
}

// Fields of the Musics.
func (Musics) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name"),
		field.
			String("description").
			MaxLen(280).
			Optional(),
	}
}

// Edges of the Musics.
func (Musics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("users", Users.Type).
			Ref("musics").
			Required(),
		edge.
			From("file", Files.Type).
			Ref("musics").
			Required().
			Unique(),
		edge.
			From("artists", Artists.Type).
			Ref("musics"),
		edge.
			From("playlists", Playlists.Type).
			Ref("musics"),
		edge.
			From("albums", Albums.Type).
			Ref("musics"),
	}
}

func (Musics) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}
