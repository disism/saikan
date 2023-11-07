package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
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

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Devices.Type),
		edge.To("playlists", Playlists.Type),
		edge.To("albums", Albums.Type),
		edge.To("musics", Musics.Type),
	}
}

func (Users) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
	}
}
