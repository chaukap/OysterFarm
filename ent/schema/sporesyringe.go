package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SporeSyringe holds the schema definition for the SporeSyringe entity.
type SporeSyringe struct {
	ent.Schema
}

// Fields of the SporeSyringe.
func (SporeSyringe) Fields() []ent.Field {
	return []ent.Field{
		field.Time("RecievedDate").
			Default(time.Now),
		field.String("Species").
			Default("unknown"),
		field.String("Supplier").
			Default("unknown"),
	}
}

// Edges of the SporeSyringe.
func (SporeSyringe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("grainJars", GrainJar.Type),
	}
}
