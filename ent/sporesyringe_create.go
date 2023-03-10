// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"example/myco-api/ent/grainjar"
	"example/myco-api/ent/sporesyringe"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SporeSyringeCreate is the builder for creating a SporeSyringe entity.
type SporeSyringeCreate struct {
	config
	mutation *SporeSyringeMutation
	hooks    []Hook
}

// SetRecievedDate sets the "RecievedDate" field.
func (ssc *SporeSyringeCreate) SetRecievedDate(t time.Time) *SporeSyringeCreate {
	ssc.mutation.SetRecievedDate(t)
	return ssc
}

// SetNillableRecievedDate sets the "RecievedDate" field if the given value is not nil.
func (ssc *SporeSyringeCreate) SetNillableRecievedDate(t *time.Time) *SporeSyringeCreate {
	if t != nil {
		ssc.SetRecievedDate(*t)
	}
	return ssc
}

// SetSpecies sets the "Species" field.
func (ssc *SporeSyringeCreate) SetSpecies(s string) *SporeSyringeCreate {
	ssc.mutation.SetSpecies(s)
	return ssc
}

// SetNillableSpecies sets the "Species" field if the given value is not nil.
func (ssc *SporeSyringeCreate) SetNillableSpecies(s *string) *SporeSyringeCreate {
	if s != nil {
		ssc.SetSpecies(*s)
	}
	return ssc
}

// SetSupplier sets the "Supplier" field.
func (ssc *SporeSyringeCreate) SetSupplier(s string) *SporeSyringeCreate {
	ssc.mutation.SetSupplier(s)
	return ssc
}

// SetNillableSupplier sets the "Supplier" field if the given value is not nil.
func (ssc *SporeSyringeCreate) SetNillableSupplier(s *string) *SporeSyringeCreate {
	if s != nil {
		ssc.SetSupplier(*s)
	}
	return ssc
}

// AddGrainJarIDs adds the "grainJars" edge to the GrainJar entity by IDs.
func (ssc *SporeSyringeCreate) AddGrainJarIDs(ids ...int) *SporeSyringeCreate {
	ssc.mutation.AddGrainJarIDs(ids...)
	return ssc
}

// AddGrainJars adds the "grainJars" edges to the GrainJar entity.
func (ssc *SporeSyringeCreate) AddGrainJars(g ...*GrainJar) *SporeSyringeCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ssc.AddGrainJarIDs(ids...)
}

// Mutation returns the SporeSyringeMutation object of the builder.
func (ssc *SporeSyringeCreate) Mutation() *SporeSyringeMutation {
	return ssc.mutation
}

// Save creates the SporeSyringe in the database.
func (ssc *SporeSyringeCreate) Save(ctx context.Context) (*SporeSyringe, error) {
	ssc.defaults()
	return withHooks[*SporeSyringe, SporeSyringeMutation](ctx, ssc.sqlSave, ssc.mutation, ssc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ssc *SporeSyringeCreate) SaveX(ctx context.Context) *SporeSyringe {
	v, err := ssc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ssc *SporeSyringeCreate) Exec(ctx context.Context) error {
	_, err := ssc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssc *SporeSyringeCreate) ExecX(ctx context.Context) {
	if err := ssc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ssc *SporeSyringeCreate) defaults() {
	if _, ok := ssc.mutation.RecievedDate(); !ok {
		v := sporesyringe.DefaultRecievedDate()
		ssc.mutation.SetRecievedDate(v)
	}
	if _, ok := ssc.mutation.Species(); !ok {
		v := sporesyringe.DefaultSpecies
		ssc.mutation.SetSpecies(v)
	}
	if _, ok := ssc.mutation.Supplier(); !ok {
		v := sporesyringe.DefaultSupplier
		ssc.mutation.SetSupplier(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ssc *SporeSyringeCreate) check() error {
	if _, ok := ssc.mutation.RecievedDate(); !ok {
		return &ValidationError{Name: "RecievedDate", err: errors.New(`ent: missing required field "SporeSyringe.RecievedDate"`)}
	}
	if _, ok := ssc.mutation.Species(); !ok {
		return &ValidationError{Name: "Species", err: errors.New(`ent: missing required field "SporeSyringe.Species"`)}
	}
	if _, ok := ssc.mutation.Supplier(); !ok {
		return &ValidationError{Name: "Supplier", err: errors.New(`ent: missing required field "SporeSyringe.Supplier"`)}
	}
	return nil
}

func (ssc *SporeSyringeCreate) sqlSave(ctx context.Context) (*SporeSyringe, error) {
	if err := ssc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ssc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ssc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ssc.mutation.id = &_node.ID
	ssc.mutation.done = true
	return _node, nil
}

func (ssc *SporeSyringeCreate) createSpec() (*SporeSyringe, *sqlgraph.CreateSpec) {
	var (
		_node = &SporeSyringe{config: ssc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sporesyringe.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sporesyringe.FieldID,
			},
		}
	)
	if value, ok := ssc.mutation.RecievedDate(); ok {
		_spec.SetField(sporesyringe.FieldRecievedDate, field.TypeTime, value)
		_node.RecievedDate = value
	}
	if value, ok := ssc.mutation.Species(); ok {
		_spec.SetField(sporesyringe.FieldSpecies, field.TypeString, value)
		_node.Species = value
	}
	if value, ok := ssc.mutation.Supplier(); ok {
		_spec.SetField(sporesyringe.FieldSupplier, field.TypeString, value)
		_node.Supplier = value
	}
	if nodes := ssc.mutation.GrainJarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sporesyringe.GrainJarsTable,
			Columns: []string{sporesyringe.GrainJarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: grainjar.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SporeSyringeCreateBulk is the builder for creating many SporeSyringe entities in bulk.
type SporeSyringeCreateBulk struct {
	config
	builders []*SporeSyringeCreate
}

// Save creates the SporeSyringe entities in the database.
func (sscb *SporeSyringeCreateBulk) Save(ctx context.Context) ([]*SporeSyringe, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sscb.builders))
	nodes := make([]*SporeSyringe, len(sscb.builders))
	mutators := make([]Mutator, len(sscb.builders))
	for i := range sscb.builders {
		func(i int, root context.Context) {
			builder := sscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SporeSyringeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sscb *SporeSyringeCreateBulk) SaveX(ctx context.Context) []*SporeSyringe {
	v, err := sscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sscb *SporeSyringeCreateBulk) Exec(ctx context.Context) error {
	_, err := sscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sscb *SporeSyringeCreateBulk) ExecX(ctx context.Context) {
	if err := sscb.Exec(ctx); err != nil {
		panic(err)
	}
}
