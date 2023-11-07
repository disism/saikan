package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Playlist holds the schema definition for the Playlist entity.
type Playlist struct {
	ent.Schema
}

// Fields of the Playlist.
func (Playlist) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.
			Bool("private").
			Default(false),
	}
}

// Edges of the Playlist.
func (Playlist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("playlists").
			Required().
			Unique(),
		edge.To("musics", Music.Type),
		edge.
			From("image", Image.Type).
			Ref("playlists").
			Unique(),
	}
}

func (Playlist) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}
