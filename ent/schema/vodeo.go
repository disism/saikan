package schema

import "entgo.io/ent"

// Vodeo holds the schema definition for the Vodeo entity.
type Vodeo struct {
	ent.Schema
}

// Fields of the Vodeo.
func (Vodeo) Fields() []ent.Field {
	return nil
}

// Edges of the Vodeo.
func (Vodeo) Edges() []ent.Edge {
	return nil
}
