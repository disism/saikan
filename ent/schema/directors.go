package schema

import "entgo.io/ent"

// Directors holds the schema definition for the Directors entity.
type Directors struct {
	ent.Schema
}

// Fields of the Directors.
func (Directors) Fields() []ent.Field {
	return nil
}

// Edges of the Directors.
func (Directors) Edges() []ent.Edge {
	return nil
}
