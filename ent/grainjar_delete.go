// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"example/myco-api/ent/grainjar"
	"example/myco-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GrainJarDelete is the builder for deleting a GrainJar entity.
type GrainJarDelete struct {
	config
	hooks    []Hook
	mutation *GrainJarMutation
}

// Where appends a list predicates to the GrainJarDelete builder.
func (gjd *GrainJarDelete) Where(ps ...predicate.GrainJar) *GrainJarDelete {
	gjd.mutation.Where(ps...)
	return gjd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gjd *GrainJarDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GrainJarMutation](ctx, gjd.sqlExec, gjd.mutation, gjd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gjd *GrainJarDelete) ExecX(ctx context.Context) int {
	n, err := gjd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gjd *GrainJarDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: grainjar.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: grainjar.FieldID,
			},
		},
	}
	if ps := gjd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gjd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gjd.mutation.done = true
	return affected, err
}

// GrainJarDeleteOne is the builder for deleting a single GrainJar entity.
type GrainJarDeleteOne struct {
	gjd *GrainJarDelete
}

// Where appends a list predicates to the GrainJarDelete builder.
func (gjdo *GrainJarDeleteOne) Where(ps ...predicate.GrainJar) *GrainJarDeleteOne {
	gjdo.gjd.mutation.Where(ps...)
	return gjdo
}

// Exec executes the deletion query.
func (gjdo *GrainJarDeleteOne) Exec(ctx context.Context) error {
	n, err := gjdo.gjd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{grainjar.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gjdo *GrainJarDeleteOne) ExecX(ctx context.Context) {
	if err := gjdo.Exec(ctx); err != nil {
		panic(err)
	}
}
