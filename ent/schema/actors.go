package schema

import "entgo.io/ent"

// Actors holds the schema definition for the Actors entity.
type Actors struct {
	ent.Schema
}

// Fields of the Actors.
func (Actors) Fields() []ent.Field {
	return nil
}

// Edges of the Actors.
func (Actors) Edges() []ent.Edge {
	return nil
}
