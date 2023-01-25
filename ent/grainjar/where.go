// Code generated by ent, DO NOT EDIT.

package grainjar

import (
	"example/myco-api/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldLTE(FieldID, id))
}

// InnoculationDate applies equality check predicate on the "InnoculationDate" field. It's identical to InnoculationDateEQ.
func InnoculationDate(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldEQ(FieldInnoculationDate, v))
}

// Grain applies equality check predicate on the "Grain" field. It's identical to GrainEQ.
func Grain(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldEQ(FieldGrain, v))
}

// HarvestDate applies equality check predicate on the "HarvestDate" field. It's identical to HarvestDateEQ.
func HarvestDate(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldEQ(FieldHarvestDate, v))
}

// InnoculationDateEQ applies the EQ predicate on the "InnoculationDate" field.
func InnoculationDateEQ(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldEQ(FieldInnoculationDate, v))
}

// InnoculationDateNEQ applies the NEQ predicate on the "InnoculationDate" field.
func InnoculationDateNEQ(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldNEQ(FieldInnoculationDate, v))
}

// InnoculationDateIn applies the In predicate on the "InnoculationDate" field.
func InnoculationDateIn(vs ...time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldIn(FieldInnoculationDate, vs...))
}

// InnoculationDateNotIn applies the NotIn predicate on the "InnoculationDate" field.
func InnoculationDateNotIn(vs ...time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldNotIn(FieldInnoculationDate, vs...))
}

// InnoculationDateGT applies the GT predicate on the "InnoculationDate" field.
func InnoculationDateGT(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldGT(FieldInnoculationDate, v))
}

// InnoculationDateGTE applies the GTE predicate on the "InnoculationDate" field.
func InnoculationDateGTE(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldGTE(FieldInnoculationDate, v))
}

// InnoculationDateLT applies the LT predicate on the "InnoculationDate" field.
func InnoculationDateLT(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldLT(FieldInnoculationDate, v))
}

// InnoculationDateLTE applies the LTE predicate on the "InnoculationDate" field.
func InnoculationDateLTE(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldLTE(FieldInnoculationDate, v))
}

// GrainEQ applies the EQ predicate on the "Grain" field.
func GrainEQ(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldEQ(FieldGrain, v))
}

// GrainNEQ applies the NEQ predicate on the "Grain" field.
func GrainNEQ(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldNEQ(FieldGrain, v))
}

// GrainIn applies the In predicate on the "Grain" field.
func GrainIn(vs ...string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldIn(FieldGrain, vs...))
}

// GrainNotIn applies the NotIn predicate on the "Grain" field.
func GrainNotIn(vs ...string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldNotIn(FieldGrain, vs...))
}

// GrainGT applies the GT predicate on the "Grain" field.
func GrainGT(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldGT(FieldGrain, v))
}

// GrainGTE applies the GTE predicate on the "Grain" field.
func GrainGTE(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldGTE(FieldGrain, v))
}

// GrainLT applies the LT predicate on the "Grain" field.
func GrainLT(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldLT(FieldGrain, v))
}

// GrainLTE applies the LTE predicate on the "Grain" field.
func GrainLTE(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldLTE(FieldGrain, v))
}

// GrainContains applies the Contains predicate on the "Grain" field.
func GrainContains(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldContains(FieldGrain, v))
}

// GrainHasPrefix applies the HasPrefix predicate on the "Grain" field.
func GrainHasPrefix(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldHasPrefix(FieldGrain, v))
}

// GrainHasSuffix applies the HasSuffix predicate on the "Grain" field.
func GrainHasSuffix(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldHasSuffix(FieldGrain, v))
}

// GrainEqualFold applies the EqualFold predicate on the "Grain" field.
func GrainEqualFold(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldEqualFold(FieldGrain, v))
}

// GrainContainsFold applies the ContainsFold predicate on the "Grain" field.
func GrainContainsFold(v string) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldContainsFold(FieldGrain, v))
}

// HarvestDateEQ applies the EQ predicate on the "HarvestDate" field.
func HarvestDateEQ(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldEQ(FieldHarvestDate, v))
}

// HarvestDateNEQ applies the NEQ predicate on the "HarvestDate" field.
func HarvestDateNEQ(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldNEQ(FieldHarvestDate, v))
}

// HarvestDateIn applies the In predicate on the "HarvestDate" field.
func HarvestDateIn(vs ...time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldIn(FieldHarvestDate, vs...))
}

// HarvestDateNotIn applies the NotIn predicate on the "HarvestDate" field.
func HarvestDateNotIn(vs ...time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldNotIn(FieldHarvestDate, vs...))
}

// HarvestDateGT applies the GT predicate on the "HarvestDate" field.
func HarvestDateGT(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldGT(FieldHarvestDate, v))
}

// HarvestDateGTE applies the GTE predicate on the "HarvestDate" field.
func HarvestDateGTE(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldGTE(FieldHarvestDate, v))
}

// HarvestDateLT applies the LT predicate on the "HarvestDate" field.
func HarvestDateLT(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldLT(FieldHarvestDate, v))
}

// HarvestDateLTE applies the LTE predicate on the "HarvestDate" field.
func HarvestDateLTE(v time.Time) predicate.GrainJar {
	return predicate.GrainJar(sql.FieldLTE(FieldHarvestDate, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GrainJar) predicate.GrainJar {
	return predicate.GrainJar(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GrainJar) predicate.GrainJar {
	return predicate.GrainJar(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GrainJar) predicate.GrainJar {
	return predicate.GrainJar(func(s *sql.Selector) {
		p(s.Not())
	})
}
