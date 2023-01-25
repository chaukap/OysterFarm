package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
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
		field.Time("HarvestDate"),
    }
}

// Edges of the GrainJar.
func (GrainJar) Edges() []ent.Edge {
	return nil
}
