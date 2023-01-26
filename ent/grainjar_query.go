// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"example/myco-api/ent/grainjar"
	"example/myco-api/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GrainJarQuery is the builder for querying GrainJar entities.
type GrainJarQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.GrainJar
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GrainJarQuery builder.
func (gjq *GrainJarQuery) Where(ps ...predicate.GrainJar) *GrainJarQuery {
	gjq.predicates = append(gjq.predicates, ps...)
	return gjq
}

// Limit the number of records to be returned by this query.
func (gjq *GrainJarQuery) Limit(limit int) *GrainJarQuery {
	gjq.ctx.Limit = &limit
	return gjq
}

// Offset to start from.
func (gjq *GrainJarQuery) Offset(offset int) *GrainJarQuery {
	gjq.ctx.Offset = &offset
	return gjq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gjq *GrainJarQuery) Unique(unique bool) *GrainJarQuery {
	gjq.ctx.Unique = &unique
	return gjq
}

// Order specifies how the records should be ordered.
func (gjq *GrainJarQuery) Order(o ...OrderFunc) *GrainJarQuery {
	gjq.order = append(gjq.order, o...)
	return gjq
}

// First returns the first GrainJar entity from the query.
// Returns a *NotFoundError when no GrainJar was found.
func (gjq *GrainJarQuery) First(ctx context.Context) (*GrainJar, error) {
	nodes, err := gjq.Limit(1).All(setContextOp(ctx, gjq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{grainjar.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gjq *GrainJarQuery) FirstX(ctx context.Context) *GrainJar {
	node, err := gjq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GrainJar ID from the query.
// Returns a *NotFoundError when no GrainJar ID was found.
func (gjq *GrainJarQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gjq.Limit(1).IDs(setContextOp(ctx, gjq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{grainjar.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gjq *GrainJarQuery) FirstIDX(ctx context.Context) int {
	id, err := gjq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GrainJar entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GrainJar entity is found.
// Returns a *NotFoundError when no GrainJar entities are found.
func (gjq *GrainJarQuery) Only(ctx context.Context) (*GrainJar, error) {
	nodes, err := gjq.Limit(2).All(setContextOp(ctx, gjq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{grainjar.Label}
	default:
		return nil, &NotSingularError{grainjar.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gjq *GrainJarQuery) OnlyX(ctx context.Context) *GrainJar {
	node, err := gjq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GrainJar ID in the query.
// Returns a *NotSingularError when more than one GrainJar ID is found.
// Returns a *NotFoundError when no entities are found.
func (gjq *GrainJarQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gjq.Limit(2).IDs(setContextOp(ctx, gjq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{grainjar.Label}
	default:
		err = &NotSingularError{grainjar.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gjq *GrainJarQuery) OnlyIDX(ctx context.Context) int {
	id, err := gjq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GrainJars.
func (gjq *GrainJarQuery) All(ctx context.Context) ([]*GrainJar, error) {
	ctx = setContextOp(ctx, gjq.ctx, "All")
	if err := gjq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GrainJar, *GrainJarQuery]()
	return withInterceptors[[]*GrainJar](ctx, gjq, qr, gjq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gjq *GrainJarQuery) AllX(ctx context.Context) []*GrainJar {
	nodes, err := gjq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GrainJar IDs.
func (gjq *GrainJarQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, gjq.ctx, "IDs")
	if err := gjq.Select(grainjar.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gjq *GrainJarQuery) IDsX(ctx context.Context) []int {
	ids, err := gjq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gjq *GrainJarQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gjq.ctx, "Count")
	if err := gjq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gjq, querierCount[*GrainJarQuery](), gjq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gjq *GrainJarQuery) CountX(ctx context.Context) int {
	count, err := gjq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gjq *GrainJarQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gjq.ctx, "Exist")
	switch _, err := gjq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gjq *GrainJarQuery) ExistX(ctx context.Context) bool {
	exist, err := gjq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GrainJarQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gjq *GrainJarQuery) Clone() *GrainJarQuery {
	if gjq == nil {
		return nil
	}
	return &GrainJarQuery{
		config:     gjq.config,
		ctx:        gjq.ctx.Clone(),
		order:      append([]OrderFunc{}, gjq.order...),
		inters:     append([]Interceptor{}, gjq.inters...),
		predicates: append([]predicate.GrainJar{}, gjq.predicates...),
		// clone intermediate query.
		sql:  gjq.sql.Clone(),
		path: gjq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		InnoculationDate time.Time `json:"InnoculationDate,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GrainJar.Query().
//		GroupBy(grainjar.FieldInnoculationDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gjq *GrainJarQuery) GroupBy(field string, fields ...string) *GrainJarGroupBy {
	gjq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GrainJarGroupBy{build: gjq}
	grbuild.flds = &gjq.ctx.Fields
	grbuild.label = grainjar.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		InnoculationDate time.Time `json:"InnoculationDate,omitempty"`
//	}
//
//	client.GrainJar.Query().
//		Select(grainjar.FieldInnoculationDate).
//		Scan(ctx, &v)
func (gjq *GrainJarQuery) Select(fields ...string) *GrainJarSelect {
	gjq.ctx.Fields = append(gjq.ctx.Fields, fields...)
	sbuild := &GrainJarSelect{GrainJarQuery: gjq}
	sbuild.label = grainjar.Label
	sbuild.flds, sbuild.scan = &gjq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GrainJarSelect configured with the given aggregations.
func (gjq *GrainJarQuery) Aggregate(fns ...AggregateFunc) *GrainJarSelect {
	return gjq.Select().Aggregate(fns...)
}

func (gjq *GrainJarQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gjq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gjq); err != nil {
				return err
			}
		}
	}
	for _, f := range gjq.ctx.Fields {
		if !grainjar.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gjq.path != nil {
		prev, err := gjq.path(ctx)
		if err != nil {
			return err
		}
		gjq.sql = prev
	}
	return nil
}

func (gjq *GrainJarQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GrainJar, error) {
	var (
		nodes   = []*GrainJar{}
		withFKs = gjq.withFKs
		_spec   = gjq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, grainjar.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GrainJar).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GrainJar{config: gjq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gjq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gjq *GrainJarQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gjq.querySpec()
	_spec.Node.Columns = gjq.ctx.Fields
	if len(gjq.ctx.Fields) > 0 {
		_spec.Unique = gjq.ctx.Unique != nil && *gjq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gjq.driver, _spec)
}

func (gjq *GrainJarQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grainjar.Table,
			Columns: grainjar.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: grainjar.FieldID,
			},
		},
		From:   gjq.sql,
		Unique: true,
	}
	if unique := gjq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gjq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grainjar.FieldID)
		for i := range fields {
			if fields[i] != grainjar.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gjq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gjq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gjq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gjq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gjq *GrainJarQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gjq.driver.Dialect())
	t1 := builder.Table(grainjar.Table)
	columns := gjq.ctx.Fields
	if len(columns) == 0 {
		columns = grainjar.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gjq.sql != nil {
		selector = gjq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gjq.ctx.Unique != nil && *gjq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gjq.predicates {
		p(selector)
	}
	for _, p := range gjq.order {
		p(selector)
	}
	if offset := gjq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gjq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GrainJarGroupBy is the group-by builder for GrainJar entities.
type GrainJarGroupBy struct {
	selector
	build *GrainJarQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gjgb *GrainJarGroupBy) Aggregate(fns ...AggregateFunc) *GrainJarGroupBy {
	gjgb.fns = append(gjgb.fns, fns...)
	return gjgb
}

// Scan applies the selector query and scans the result into the given value.
func (gjgb *GrainJarGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gjgb.build.ctx, "GroupBy")
	if err := gjgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GrainJarQuery, *GrainJarGroupBy](ctx, gjgb.build, gjgb, gjgb.build.inters, v)
}

func (gjgb *GrainJarGroupBy) sqlScan(ctx context.Context, root *GrainJarQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gjgb.fns))
	for _, fn := range gjgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gjgb.flds)+len(gjgb.fns))
		for _, f := range *gjgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gjgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gjgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GrainJarSelect is the builder for selecting fields of GrainJar entities.
type GrainJarSelect struct {
	*GrainJarQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gjs *GrainJarSelect) Aggregate(fns ...AggregateFunc) *GrainJarSelect {
	gjs.fns = append(gjs.fns, fns...)
	return gjs
}

// Scan applies the selector query and scans the result into the given value.
func (gjs *GrainJarSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gjs.ctx, "Select")
	if err := gjs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GrainJarQuery, *GrainJarSelect](ctx, gjs.GrainJarQuery, gjs, gjs.inters, v)
}

func (gjs *GrainJarSelect) sqlScan(ctx context.Context, root *GrainJarQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gjs.fns))
	for _, fn := range gjs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gjs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gjs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
