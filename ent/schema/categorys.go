package schema

import "entgo.io/ent"

// Categorys holds the schema definition for the Categorys entity.
type Categorys struct {
	ent.Schema
}

// Fields of the Categorys.
func (Categorys) Fields() []ent.Field {
	return nil
}

// Edges of the Categorys.
func (Categorys) Edges() []ent.Edge {
	return nil
}
