// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Naist4869/awesomeProject1/ent/migrate"

	"github.com/Naist4869/awesomeProject1/ent/issuex1355"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// IssueX1355 is the client for interacting with the IssueX1355 builders.
	IssueX1355 *IssueX1355Client
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.IssueX1355 = NewIssueX1355Client(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		IssueX1355: NewIssueX1355Client(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		config:     cfg,
		IssueX1355: NewIssueX1355Client(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		IssueX1355.
//		Query().
//		Count(ctx)
//
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
	c.IssueX1355.Use(hooks...)
}

// IssueX1355Client is a client for the IssueX1355 schema.
type IssueX1355Client struct {
	config
}

// NewIssueX1355Client returns a client for the IssueX1355 from the given config.
func NewIssueX1355Client(c config) *IssueX1355Client {
	return &IssueX1355Client{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `issuex1355.Hooks(f(g(h())))`.
func (c *IssueX1355Client) Use(hooks ...Hook) {
	c.hooks.IssueX1355 = append(c.hooks.IssueX1355, hooks...)
}

// Create returns a create builder for IssueX1355.
func (c *IssueX1355Client) Create() *IssueX1355Create {
	mutation := newIssueX1355Mutation(c.config, OpCreate)
	return &IssueX1355Create{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of IssueX1355 entities.
func (c *IssueX1355Client) CreateBulk(builders ...*IssueX1355Create) *IssueX1355CreateBulk {
	return &IssueX1355CreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for IssueX1355.
func (c *IssueX1355Client) Update() *IssueX1355Update {
	mutation := newIssueX1355Mutation(c.config, OpUpdate)
	return &IssueX1355Update{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IssueX1355Client) UpdateOne(i *IssueX1355) *IssueX1355UpdateOne {
	mutation := newIssueX1355Mutation(c.config, OpUpdateOne, withIssueX1355(i))
	return &IssueX1355UpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IssueX1355Client) UpdateOneID(id int) *IssueX1355UpdateOne {
	mutation := newIssueX1355Mutation(c.config, OpUpdateOne, withIssueX1355ID(id))
	return &IssueX1355UpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for IssueX1355.
func (c *IssueX1355Client) Delete() *IssueX1355Delete {
	mutation := newIssueX1355Mutation(c.config, OpDelete)
	return &IssueX1355Delete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *IssueX1355Client) DeleteOne(i *IssueX1355) *IssueX1355DeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *IssueX1355Client) DeleteOneID(id int) *IssueX1355DeleteOne {
	builder := c.Delete().Where(issuex1355.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IssueX1355DeleteOne{builder}
}

// Query returns a query builder for IssueX1355.
func (c *IssueX1355Client) Query() *IssueX1355Query {
	return &IssueX1355Query{config: c.config}
}

// Get returns a IssueX1355 entity by its id.
func (c *IssueX1355Client) Get(ctx context.Context, id int) (*IssueX1355, error) {
	return c.Query().Where(issuex1355.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IssueX1355Client) GetX(ctx context.Context, id int) *IssueX1355 {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *IssueX1355Client) Hooks() []Hook {
	return c.hooks.IssueX1355
}
