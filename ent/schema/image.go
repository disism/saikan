package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.
			Int32("width").
			Comment("IMAGE WIDTH"),
		field.
			Int32("height").
			Comment("IMAGE HEIGHT"),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("file", File.Type).
			Ref("images").
			Required().
			Unique(),
		edge.To("albums", Album.Type),
		edge.To("playlists", Playlist.Type),
	}
}

func (Image) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}
