package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Music holds the schema definition for the Music entity.
type Music struct {
	ent.Schema
}

// Fields of the Music.
func (Music) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name"),
		field.
			String("description").
			MaxLen(280).
			Optional(),
	}
}

// Edges of the Music.
func (Music) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("musics").
			Required(),
		edge.From("file", File.Type).
			Ref("musics").
			Required().
			Unique(),
		edge.From("artists", Artist.Type).
			Ref("musics"),
		edge.From("playlists", Playlist.Type).
			Ref("musics"),
		edge.From("albums", Album.Type).
			Ref("musics"),
	}
}

func (Music) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}
