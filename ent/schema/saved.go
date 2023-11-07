package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Saved holds the schema definition for the Saved entity.
type Saved struct {
	ent.Schema
}

// Fields of the Saved.
func (Saved) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(255),
		field.
			String("caption").
			MaxLen(255).
			Optional().
			Comment("the descriptive text or title of a document, image, or other media element. It is used to provide a short description of the content, characteristics or context of a document."),
	}
}

// Edges of the Saved.
func (Saved) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("file", File.Type).
			Ref("saves").
			Unique(),
		edge.
			From("owner", User.Type).
			Ref("saves").
			Unique(),
		edge.
			From("dir", Dir.Type).
			Ref("saves"),
	}
}

func (Saved) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}
