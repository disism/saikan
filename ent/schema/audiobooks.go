package schema

import "entgo.io/ent"

// Audiobooks holds the schema definition for the Audiobooks entity.
type Audiobooks struct {
	ent.Schema
}

// Fields of the Audiobooks.
func (Audiobooks) Fields() []ent.Field {
	return nil
}

// Edges of the Audiobooks.
func (Audiobooks) Edges() []ent.Edge {
	return nil
}
