// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"example/myco-api/ent/migrate"

	"example/myco-api/ent/grainjar"
	"example/myco-api/ent/sporesyringe"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// GrainJar is the client for interacting with the GrainJar builders.
	GrainJar *GrainJarClient
	// SporeSyringe is the client for interacting with the SporeSyringe builders.
	SporeSyringe *SporeSyringeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.GrainJar = NewGrainJarClient(c.config)
	c.SporeSyringe = NewSporeSyringeClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		GrainJar:     NewGrainJarClient(cfg),
		SporeSyringe: NewSporeSyringeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		GrainJar:     NewGrainJarClient(cfg),
		SporeSyringe: NewSporeSyringeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		GrainJar.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.GrainJar.Use(hooks...)
	c.SporeSyringe.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.GrainJar.Intercept(interceptors...)
	c.SporeSyringe.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *GrainJarMutation:
		return c.GrainJar.mutate(ctx, m)
	case *SporeSyringeMutation:
		return c.SporeSyringe.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// GrainJarClient is a client for the GrainJar schema.
type GrainJarClient struct {
	config
}

// NewGrainJarClient returns a client for the GrainJar from the given config.
func NewGrainJarClient(c config) *GrainJarClient {
	return &GrainJarClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `grainjar.Hooks(f(g(h())))`.
func (c *GrainJarClient) Use(hooks ...Hook) {
	c.hooks.GrainJar = append(c.hooks.GrainJar, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `grainjar.Intercept(f(g(h())))`.
func (c *GrainJarClient) Intercept(interceptors ...Interceptor) {
	c.inters.GrainJar = append(c.inters.GrainJar, interceptors...)
}

// Create returns a builder for creating a GrainJar entity.
func (c *GrainJarClient) Create() *GrainJarCreate {
	mutation := newGrainJarMutation(c.config, OpCreate)
	return &GrainJarCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GrainJar entities.
func (c *GrainJarClient) CreateBulk(builders ...*GrainJarCreate) *GrainJarCreateBulk {
	return &GrainJarCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GrainJar.
func (c *GrainJarClient) Update() *GrainJarUpdate {
	mutation := newGrainJarMutation(c.config, OpUpdate)
	return &GrainJarUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GrainJarClient) UpdateOne(gj *GrainJar) *GrainJarUpdateOne {
	mutation := newGrainJarMutation(c.config, OpUpdateOne, withGrainJar(gj))
	return &GrainJarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GrainJarClient) UpdateOneID(id int) *GrainJarUpdateOne {
	mutation := newGrainJarMutation(c.config, OpUpdateOne, withGrainJarID(id))
	return &GrainJarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GrainJar.
func (c *GrainJarClient) Delete() *GrainJarDelete {
	mutation := newGrainJarMutation(c.config, OpDelete)
	return &GrainJarDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GrainJarClient) DeleteOne(gj *GrainJar) *GrainJarDeleteOne {
	return c.DeleteOneID(gj.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GrainJarClient) DeleteOneID(id int) *GrainJarDeleteOne {
	builder := c.Delete().Where(grainjar.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GrainJarDeleteOne{builder}
}

// Query returns a query builder for GrainJar.
func (c *GrainJarClient) Query() *GrainJarQuery {
	return &GrainJarQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGrainJar},
		inters: c.Interceptors(),
	}
}

// Get returns a GrainJar entity by its id.
func (c *GrainJarClient) Get(ctx context.Context, id int) (*GrainJar, error) {
	return c.Query().Where(grainjar.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GrainJarClient) GetX(ctx context.Context, id int) *GrainJar {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GrainJarClient) Hooks() []Hook {
	return c.hooks.GrainJar
}

// Interceptors returns the client interceptors.
func (c *GrainJarClient) Interceptors() []Interceptor {
	return c.inters.GrainJar
}

func (c *GrainJarClient) mutate(ctx context.Context, m *GrainJarMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GrainJarCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GrainJarUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GrainJarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GrainJarDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown GrainJar mutation op: %q", m.Op())
	}
}

// SporeSyringeClient is a client for the SporeSyringe schema.
type SporeSyringeClient struct {
	config
}

// NewSporeSyringeClient returns a client for the SporeSyringe from the given config.
func NewSporeSyringeClient(c config) *SporeSyringeClient {
	return &SporeSyringeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sporesyringe.Hooks(f(g(h())))`.
func (c *SporeSyringeClient) Use(hooks ...Hook) {
	c.hooks.SporeSyringe = append(c.hooks.SporeSyringe, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `sporesyringe.Intercept(f(g(h())))`.
func (c *SporeSyringeClient) Intercept(interceptors ...Interceptor) {
	c.inters.SporeSyringe = append(c.inters.SporeSyringe, interceptors...)
}

// Create returns a builder for creating a SporeSyringe entity.
func (c *SporeSyringeClient) Create() *SporeSyringeCreate {
	mutation := newSporeSyringeMutation(c.config, OpCreate)
	return &SporeSyringeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SporeSyringe entities.
func (c *SporeSyringeClient) CreateBulk(builders ...*SporeSyringeCreate) *SporeSyringeCreateBulk {
	return &SporeSyringeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SporeSyringe.
func (c *SporeSyringeClient) Update() *SporeSyringeUpdate {
	mutation := newSporeSyringeMutation(c.config, OpUpdate)
	return &SporeSyringeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SporeSyringeClient) UpdateOne(ss *SporeSyringe) *SporeSyringeUpdateOne {
	mutation := newSporeSyringeMutation(c.config, OpUpdateOne, withSporeSyringe(ss))
	return &SporeSyringeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SporeSyringeClient) UpdateOneID(id int) *SporeSyringeUpdateOne {
	mutation := newSporeSyringeMutation(c.config, OpUpdateOne, withSporeSyringeID(id))
	return &SporeSyringeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SporeSyringe.
func (c *SporeSyringeClient) Delete() *SporeSyringeDelete {
	mutation := newSporeSyringeMutation(c.config, OpDelete)
	return &SporeSyringeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SporeSyringeClient) DeleteOne(ss *SporeSyringe) *SporeSyringeDeleteOne {
	return c.DeleteOneID(ss.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SporeSyringeClient) DeleteOneID(id int) *SporeSyringeDeleteOne {
	builder := c.Delete().Where(sporesyringe.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SporeSyringeDeleteOne{builder}
}

// Query returns a query builder for SporeSyringe.
func (c *SporeSyringeClient) Query() *SporeSyringeQuery {
	return &SporeSyringeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSporeSyringe},
		inters: c.Interceptors(),
	}
}

// Get returns a SporeSyringe entity by its id.
func (c *SporeSyringeClient) Get(ctx context.Context, id int) (*SporeSyringe, error) {
	return c.Query().Where(sporesyringe.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SporeSyringeClient) GetX(ctx context.Context, id int) *SporeSyringe {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGrainJar queries the grainJar edge of a SporeSyringe.
func (c *SporeSyringeClient) QueryGrainJar(ss *SporeSyringe) *GrainJarQuery {
	query := (&GrainJarClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ss.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(sporesyringe.Table, sporesyringe.FieldID, id),
			sqlgraph.To(grainjar.Table, grainjar.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sporesyringe.GrainJarTable, sporesyringe.GrainJarColumn),
		)
		fromV = sqlgraph.Neighbors(ss.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SporeSyringeClient) Hooks() []Hook {
	return c.hooks.SporeSyringe
}

// Interceptors returns the client interceptors.
func (c *SporeSyringeClient) Interceptors() []Interceptor {
	return c.inters.SporeSyringe
}

func (c *SporeSyringeClient) mutate(ctx context.Context, m *SporeSyringeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SporeSyringeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SporeSyringeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SporeSyringeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SporeSyringeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown SporeSyringe mutation op: %q", m.Op())
	}
}
