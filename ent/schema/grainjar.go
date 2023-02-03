package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GrainJar holds the schema definition for the GrainJar entity.
type GrainJar struct {
	ent.Schema
}

// Fields of the GrainJar.
func (GrainJar) Fields() []ent.Field {
	return []ent.Field{
		field.Time("InnoculationDate").
			Default(time.Now),
		field.String("Grain").
			Default("unknown"),
		field.Time("HarvestDate").
			Optional(),
	}
}

// Edges of the GrainJar.
func (GrainJar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sporeSyringes", SporeSyringe.Type).Ref("grainJars").Unique(),
	}
}
