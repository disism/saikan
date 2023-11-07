package schema

import "entgo.io/ent"

// Actor holds the schema definition for the Actor entity.
type Actor struct {
	ent.Schema
}

// Fields of the Actor.
func (Actor) Fields() []ent.Field {
	return nil
}

// Edges of the Actor.
func (Actor) Edges() []ent.Edge {
	return nil
}
