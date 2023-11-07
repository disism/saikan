package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("username").
			MaxLen(255).
			MinLen(3).
			Unique(),
		field.
			String("password").
			Optional(),
		field.
			String("email").
			Unique().
			Optional(),
		field.
			String("name").
			Optional(),
		field.
			String("bio").
			Optional(),
		field.
			String("avatar").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
		edge.To("playlists", Playlist.Type),
		edge.To("albums", Album.Type),
		edge.To("musics", Music.Type),
		edge.To("dirs", Dir.Type),
		edge.To("saves", Saved.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}
