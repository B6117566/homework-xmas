package schema

import "github.com/facebookincubator/ent"

// DepenseMedicine holds the schema definition for the DepenseMedicine entity.
type DepenseMedicine struct {
	ent.Schema
}

// Fields of the DepenseMedicine.
func (DepenseMedicine) Fields() []ent.Field {
	return nil
}

// Edges of the DepenseMedicine.
func (DepenseMedicine) Edges() []ent.Edge {
	return nil
}
