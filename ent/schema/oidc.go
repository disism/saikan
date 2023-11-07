package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Oidc holds the schema definition for the Oidc entity.
type Oidc struct {
	ent.Schema
}

// Fields of the Oidc.
func (Oidc) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			Unique(),
		field.
			String("configuration_endpoint"),
	}
}

// Edges of the Oidc.
func (Oidc) Edges() []ent.Edge {
	return nil
}

func (Oidc) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
	}
}
