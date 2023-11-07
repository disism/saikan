package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Playlists holds the schema definition for the Playlists entity.
type Playlists struct {
	ent.Schema
}

// Fields of the Playlists.
func (Playlists) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			NotEmpty(),
		field.
			String("description").
			Optional(),
		field.
			Bool("private").
			Default(false),
	}
}

// Edges of the Playlists.
func (Playlists) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("owner", Users.Type).
			Ref("playlists").
			Required().
			Unique(),
		edge.
			To("musics", Musics.Type),
		edge.
			From("image", Images.Type).
			Ref("playlists").
			Unique(),
	}
}

func (Playlists) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}
