// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Yostardev/errors/store/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/Yostardev/errors/store/ent/errorcode"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ErrorCode is the client for interacting with the ErrorCode builders.
	ErrorCode *ErrorCodeClient
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
	c.ErrorCode = NewErrorCodeClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		ctx:       ctx,
		config:    cfg,
		ErrorCode: NewErrorCodeClient(cfg),
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
		ctx:       ctx,
		config:    cfg,
		ErrorCode: NewErrorCodeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ErrorCode.
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
	c.ErrorCode.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.ErrorCode.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ErrorCodeMutation:
		return c.ErrorCode.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ErrorCodeClient is a client for the ErrorCode schema.
type ErrorCodeClient struct {
	config
}

// NewErrorCodeClient returns a client for the ErrorCode from the given config.
func NewErrorCodeClient(c config) *ErrorCodeClient {
	return &ErrorCodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `errorcode.Hooks(f(g(h())))`.
func (c *ErrorCodeClient) Use(hooks ...Hook) {
	c.hooks.ErrorCode = append(c.hooks.ErrorCode, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `errorcode.Intercept(f(g(h())))`.
func (c *ErrorCodeClient) Intercept(interceptors ...Interceptor) {
	c.inters.ErrorCode = append(c.inters.ErrorCode, interceptors...)
}

// Create returns a builder for creating a ErrorCode entity.
func (c *ErrorCodeClient) Create() *ErrorCodeCreate {
	mutation := newErrorCodeMutation(c.config, OpCreate)
	return &ErrorCodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ErrorCode entities.
func (c *ErrorCodeClient) CreateBulk(builders ...*ErrorCodeCreate) *ErrorCodeCreateBulk {
	return &ErrorCodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ErrorCode.
func (c *ErrorCodeClient) Update() *ErrorCodeUpdate {
	mutation := newErrorCodeMutation(c.config, OpUpdate)
	return &ErrorCodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ErrorCodeClient) UpdateOne(ec *ErrorCode) *ErrorCodeUpdateOne {
	mutation := newErrorCodeMutation(c.config, OpUpdateOne, withErrorCode(ec))
	return &ErrorCodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ErrorCodeClient) UpdateOneID(id int64) *ErrorCodeUpdateOne {
	mutation := newErrorCodeMutation(c.config, OpUpdateOne, withErrorCodeID(id))
	return &ErrorCodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ErrorCode.
func (c *ErrorCodeClient) Delete() *ErrorCodeDelete {
	mutation := newErrorCodeMutation(c.config, OpDelete)
	return &ErrorCodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ErrorCodeClient) DeleteOne(ec *ErrorCode) *ErrorCodeDeleteOne {
	return c.DeleteOneID(ec.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ErrorCodeClient) DeleteOneID(id int64) *ErrorCodeDeleteOne {
	builder := c.Delete().Where(errorcode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ErrorCodeDeleteOne{builder}
}

// Query returns a query builder for ErrorCode.
func (c *ErrorCodeClient) Query() *ErrorCodeQuery {
	return &ErrorCodeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeErrorCode},
		inters: c.Interceptors(),
	}
}

// Get returns a ErrorCode entity by its id.
func (c *ErrorCodeClient) Get(ctx context.Context, id int64) (*ErrorCode, error) {
	return c.Query().Where(errorcode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ErrorCodeClient) GetX(ctx context.Context, id int64) *ErrorCode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ErrorCodeClient) Hooks() []Hook {
	return c.hooks.ErrorCode
}

// Interceptors returns the client interceptors.
func (c *ErrorCodeClient) Interceptors() []Interceptor {
	return c.inters.ErrorCode
}

func (c *ErrorCodeClient) mutate(ctx context.Context, m *ErrorCodeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ErrorCodeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ErrorCodeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ErrorCodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ErrorCodeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ErrorCode mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ErrorCode []ent.Hook
	}
	inters struct {
		ErrorCode []ent.Interceptor
	}
)
