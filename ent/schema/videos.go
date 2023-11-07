package schema

import "entgo.io/ent"

// Videos holds the schema definition for the Videos entity.
type Videos struct {
	ent.Schema
}

// Fields of the Videos.
func (Videos) Fields() []ent.Field {
	return nil
}

// Edges of the Videos.
func (Videos) Edges() []ent.Edge {
	return nil
}
