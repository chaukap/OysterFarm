// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GrainJarsColumns holds the columns for the "grain_jars" table.
	GrainJarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "innoculation_date", Type: field.TypeTime},
		{Name: "grain", Type: field.TypeString, Default: "unknown"},
		{Name: "harvest_date", Type: field.TypeTime, Nullable: true},
		{Name: "spore_syringe_grain_jars", Type: field.TypeInt, Nullable: true},
	}
	// GrainJarsTable holds the schema information for the "grain_jars" table.
	GrainJarsTable = &schema.Table{
		Name:       "grain_jars",
		Columns:    GrainJarsColumns,
		PrimaryKey: []*schema.Column{GrainJarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "grain_jars_spore_syringes_grainJars",
				Columns:    []*schema.Column{GrainJarsColumns[4]},
				RefColumns: []*schema.Column{SporeSyringesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SporeSyringesColumns holds the columns for the "spore_syringes" table.
	SporeSyringesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "recieved_date", Type: field.TypeTime},
		{Name: "species", Type: field.TypeString, Default: "unknown"},
		{Name: "supplier", Type: field.TypeString, Default: "unknown"},
	}
	// SporeSyringesTable holds the schema information for the "spore_syringes" table.
	SporeSyringesTable = &schema.Table{
		Name:       "spore_syringes",
		Columns:    SporeSyringesColumns,
		PrimaryKey: []*schema.Column{SporeSyringesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GrainJarsTable,
		SporeSyringesTable,
	}
)

func init() {
	GrainJarsTable.ForeignKeys[0].RefTable = SporeSyringesTable
}
