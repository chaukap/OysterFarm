// Code generated by ent, DO NOT EDIT.

package sporesyringe

import (
	"time"
)

const (
	// Label holds the string label denoting the sporesyringe type in the database.
	Label = "spore_syringe"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRecievedDate holds the string denoting the recieveddate field in the database.
	FieldRecievedDate = "recieved_date"
	// FieldSpecies holds the string denoting the species field in the database.
	FieldSpecies = "species"
	// FieldSupplier holds the string denoting the supplier field in the database.
	FieldSupplier = "supplier"
	// Table holds the table name of the sporesyringe in the database.
	Table = "spore_syringes"
)

// Columns holds all SQL columns for sporesyringe fields.
var Columns = []string{
	FieldID,
	FieldRecievedDate,
	FieldSpecies,
	FieldSupplier,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "spore_syringes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"grain_jar_spore_syringe",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRecievedDate holds the default value on creation for the "RecievedDate" field.
	DefaultRecievedDate func() time.Time
	// DefaultSpecies holds the default value on creation for the "Species" field.
	DefaultSpecies string
	// DefaultSupplier holds the default value on creation for the "Supplier" field.
	DefaultSupplier string
)