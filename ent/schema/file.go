package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("hash").
			Unique().
			Comment("ipfs hash"),
		field.
			String("name").
			NotEmpty().
			Comment("file name"),
		field.
			Uint64("size").
			Comment("file size, number of bytes in the stored file"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("musics", Music.Type),
		edge.To("images", Image.Type),
		edge.To("saves", Saved.Type),
	}
}

func (File) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("hash").
			Unique(),
	}
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
