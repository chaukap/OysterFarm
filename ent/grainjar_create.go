// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"example/myco-api/ent/grainjar"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GrainJarCreate is the builder for creating a GrainJar entity.
type GrainJarCreate struct {
	config
	mutation *GrainJarMutation
	hooks    []Hook
}

// SetInnoculationDate sets the "InnoculationDate" field.
func (gjc *GrainJarCreate) SetInnoculationDate(t time.Time) *GrainJarCreate {
	gjc.mutation.SetInnoculationDate(t)
	return gjc
}

// SetNillableInnoculationDate sets the "InnoculationDate" field if the given value is not nil.
func (gjc *GrainJarCreate) SetNillableInnoculationDate(t *time.Time) *GrainJarCreate {
	if t != nil {
		gjc.SetInnoculationDate(*t)
	}
	return gjc
}

// SetGrain sets the "Grain" field.
func (gjc *GrainJarCreate) SetGrain(s string) *GrainJarCreate {
	gjc.mutation.SetGrain(s)
	return gjc
}

// SetNillableGrain sets the "Grain" field if the given value is not nil.
func (gjc *GrainJarCreate) SetNillableGrain(s *string) *GrainJarCreate {
	if s != nil {
		gjc.SetGrain(*s)
	}
	return gjc
}

// SetHarvestDate sets the "HarvestDate" field.
func (gjc *GrainJarCreate) SetHarvestDate(t time.Time) *GrainJarCreate {
	gjc.mutation.SetHarvestDate(t)
	return gjc
}

// Mutation returns the GrainJarMutation object of the builder.
func (gjc *GrainJarCreate) Mutation() *GrainJarMutation {
	return gjc.mutation
}

// Save creates the GrainJar in the database.
func (gjc *GrainJarCreate) Save(ctx context.Context) (*GrainJar, error) {
	gjc.defaults()
	return withHooks[*GrainJar, GrainJarMutation](ctx, gjc.sqlSave, gjc.mutation, gjc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gjc *GrainJarCreate) SaveX(ctx context.Context) *GrainJar {
	v, err := gjc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gjc *GrainJarCreate) Exec(ctx context.Context) error {
	_, err := gjc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gjc *GrainJarCreate) ExecX(ctx context.Context) {
	if err := gjc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gjc *GrainJarCreate) defaults() {
	if _, ok := gjc.mutation.InnoculationDate(); !ok {
		v := grainjar.DefaultInnoculationDate()
		gjc.mutation.SetInnoculationDate(v)
	}
	if _, ok := gjc.mutation.Grain(); !ok {
		v := grainjar.DefaultGrain
		gjc.mutation.SetGrain(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gjc *GrainJarCreate) check() error {
	if _, ok := gjc.mutation.InnoculationDate(); !ok {
		return &ValidationError{Name: "InnoculationDate", err: errors.New(`ent: missing required field "GrainJar.InnoculationDate"`)}
	}
	if _, ok := gjc.mutation.Grain(); !ok {
		return &ValidationError{Name: "Grain", err: errors.New(`ent: missing required field "GrainJar.Grain"`)}
	}
	if _, ok := gjc.mutation.HarvestDate(); !ok {
		return &ValidationError{Name: "HarvestDate", err: errors.New(`ent: missing required field "GrainJar.HarvestDate"`)}
	}
	return nil
}

func (gjc *GrainJarCreate) sqlSave(ctx context.Context) (*GrainJar, error) {
	if err := gjc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gjc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gjc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gjc.mutation.id = &_node.ID
	gjc.mutation.done = true
	return _node, nil
}

func (gjc *GrainJarCreate) createSpec() (*GrainJar, *sqlgraph.CreateSpec) {
	var (
		_node = &GrainJar{config: gjc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: grainjar.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: grainjar.FieldID,
			},
		}
	)
	if value, ok := gjc.mutation.InnoculationDate(); ok {
		_spec.SetField(grainjar.FieldInnoculationDate, field.TypeTime, value)
		_node.InnoculationDate = value
	}
	if value, ok := gjc.mutation.Grain(); ok {
		_spec.SetField(grainjar.FieldGrain, field.TypeString, value)
		_node.Grain = value
	}
	if value, ok := gjc.mutation.HarvestDate(); ok {
		_spec.SetField(grainjar.FieldHarvestDate, field.TypeTime, value)
		_node.HarvestDate = value
	}
	return _node, _spec
}

// GrainJarCreateBulk is the builder for creating many GrainJar entities in bulk.
type GrainJarCreateBulk struct {
	config
	builders []*GrainJarCreate
}

// Save creates the GrainJar entities in the database.
func (gjcb *GrainJarCreateBulk) Save(ctx context.Context) ([]*GrainJar, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gjcb.builders))
	nodes := make([]*GrainJar, len(gjcb.builders))
	mutators := make([]Mutator, len(gjcb.builders))
	for i := range gjcb.builders {
		func(i int, root context.Context) {
			builder := gjcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GrainJarMutation)
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
					_, err = mutators[i+1].Mutate(root, gjcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gjcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gjcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gjcb *GrainJarCreateBulk) SaveX(ctx context.Context) []*GrainJar {
	v, err := gjcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gjcb *GrainJarCreateBulk) Exec(ctx context.Context) error {
	_, err := gjcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gjcb *GrainJarCreateBulk) ExecX(ctx context.Context) {
	if err := gjcb.Exec(ctx); err != nil {
		panic(err)
	}
}
