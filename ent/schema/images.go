package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Images holds the schema definition for the Images entity.
type Images struct {
	ent.Schema
}

// Fields of the Images.
func (Images) Fields() []ent.Field {
	return []ent.Field{
		field.
			Int32("width").
			Comment("IMAGE WIDTH"),
		field.
			Int32("height").
			Comment("IMAGE HEIGHT"),
	}
}

// Edges of the Images.
func (Images) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("file", Files.Type).
			Ref("images").
			Required().
			Unique(),
		edge.
			To("albums", Albums.Type),
		edge.
			To("playlists", Playlists.Type),
	}
}

func (Images) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}
