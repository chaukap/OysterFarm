// Code generated by ent, DO NOT EDIT.

package ent

import (
	"example/myco-api/ent/sporesyringe"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// SporeSyringe is the model entity for the SporeSyringe schema.
type SporeSyringe struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RecievedDate holds the value of the "RecievedDate" field.
	RecievedDate time.Time `json:"RecievedDate,omitempty"`
	// Species holds the value of the "Species" field.
	Species string `json:"Species,omitempty"`
	// Supplier holds the value of the "Supplier" field.
	Supplier string `json:"Supplier,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SporeSyringeQuery when eager-loading is set.
	Edges SporeSyringeEdges `json:"edges"`
}

// SporeSyringeEdges holds the relations/edges for other nodes in the graph.
type SporeSyringeEdges struct {
	// GrainJar holds the value of the grainJar edge.
	GrainJar []*GrainJar `json:"grainJar,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GrainJarOrErr returns the GrainJar value or an error if the edge
// was not loaded in eager-loading.
func (e SporeSyringeEdges) GrainJarOrErr() ([]*GrainJar, error) {
	if e.loadedTypes[0] {
		return e.GrainJar, nil
	}
	return nil, &NotLoadedError{edge: "grainJar"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SporeSyringe) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sporesyringe.FieldID:
			values[i] = new(sql.NullInt64)
		case sporesyringe.FieldSpecies, sporesyringe.FieldSupplier:
			values[i] = new(sql.NullString)
		case sporesyringe.FieldRecievedDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SporeSyringe", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SporeSyringe fields.
func (ss *SporeSyringe) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sporesyringe.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ss.ID = int(value.Int64)
		case sporesyringe.FieldRecievedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field RecievedDate", values[i])
			} else if value.Valid {
				ss.RecievedDate = value.Time
			}
		case sporesyringe.FieldSpecies:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Species", values[i])
			} else if value.Valid {
				ss.Species = value.String
			}
		case sporesyringe.FieldSupplier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Supplier", values[i])
			} else if value.Valid {
				ss.Supplier = value.String
			}
		}
	}
	return nil
}

// QueryGrainJar queries the "grainJar" edge of the SporeSyringe entity.
func (ss *SporeSyringe) QueryGrainJar() *GrainJarQuery {
	return NewSporeSyringeClient(ss.config).QueryGrainJar(ss)
}

// Update returns a builder for updating this SporeSyringe.
// Note that you need to call SporeSyringe.Unwrap() before calling this method if this SporeSyringe
// was returned from a transaction, and the transaction was committed or rolled back.
func (ss *SporeSyringe) Update() *SporeSyringeUpdateOne {
	return NewSporeSyringeClient(ss.config).UpdateOne(ss)
}

// Unwrap unwraps the SporeSyringe entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ss *SporeSyringe) Unwrap() *SporeSyringe {
	_tx, ok := ss.config.driver.(*txDriver)
	if !ok {
		panic("ent: SporeSyringe is not a transactional entity")
	}
	ss.config.driver = _tx.drv
	return ss
}

// String implements the fmt.Stringer.
func (ss *SporeSyringe) String() string {
	var builder strings.Builder
	builder.WriteString("SporeSyringe(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ss.ID))
	builder.WriteString("RecievedDate=")
	builder.WriteString(ss.RecievedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("Species=")
	builder.WriteString(ss.Species)
	builder.WriteString(", ")
	builder.WriteString("Supplier=")
	builder.WriteString(ss.Supplier)
	builder.WriteByte(')')
	return builder.String()
}

// SporeSyringes is a parsable slice of SporeSyringe.
type SporeSyringes []*SporeSyringe

func (ss SporeSyringes) config(cfg config) {
	for _i := range ss {
		ss[_i].config = cfg
	}
}
