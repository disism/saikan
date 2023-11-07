package schema

import "entgo.io/ent"

// Director holds the schema definition for the Director entity.
type Director struct {
	ent.Schema
}

// Fields of the Director.
func (Director) Fields() []ent.Field {
	return nil
}

// Edges of the Director.
func (Director) Edges() []ent.Edge {
	return nil
}
