// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"database-concurrency/ent/migrate"

	"database-concurrency/ent/createticketlog"
	"database-concurrency/ent/serviceprodiver"
	"database-concurrency/ent/ticket"
	"database-concurrency/ent/ticketevent"
	"database-concurrency/ent/transaction"
	"database-concurrency/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CreateTicketLog is the client for interacting with the CreateTicketLog builders.
	CreateTicketLog *CreateTicketLogClient
	// ServiceProdiver is the client for interacting with the ServiceProdiver builders.
	ServiceProdiver *ServiceProdiverClient
	// Ticket is the client for interacting with the Ticket builders.
	Ticket *TicketClient
	// TicketEvent is the client for interacting with the TicketEvent builders.
	TicketEvent *TicketEventClient
	// Transaction is the client for interacting with the Transaction builders.
	Transaction *TransactionClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.CreateTicketLog = NewCreateTicketLogClient(c.config)
	c.ServiceProdiver = NewServiceProdiverClient(c.config)
	c.Ticket = NewTicketClient(c.config)
	c.TicketEvent = NewTicketEventClient(c.config)
	c.Transaction = NewTransactionClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:             ctx,
		config:          cfg,
		CreateTicketLog: NewCreateTicketLogClient(cfg),
		ServiceProdiver: NewServiceProdiverClient(cfg),
		Ticket:          NewTicketClient(cfg),
		TicketEvent:     NewTicketEventClient(cfg),
		Transaction:     NewTransactionClient(cfg),
		User:            NewUserClient(cfg),
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
		ctx:             ctx,
		config:          cfg,
		CreateTicketLog: NewCreateTicketLogClient(cfg),
		ServiceProdiver: NewServiceProdiverClient(cfg),
		Ticket:          NewTicketClient(cfg),
		TicketEvent:     NewTicketEventClient(cfg),
		Transaction:     NewTransactionClient(cfg),
		User:            NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CreateTicketLog.
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
	c.CreateTicketLog.Use(hooks...)
	c.ServiceProdiver.Use(hooks...)
	c.Ticket.Use(hooks...)
	c.TicketEvent.Use(hooks...)
	c.Transaction.Use(hooks...)
	c.User.Use(hooks...)
}

// CreateTicketLogClient is a client for the CreateTicketLog schema.
type CreateTicketLogClient struct {
	config
}

// NewCreateTicketLogClient returns a client for the CreateTicketLog from the given config.
func NewCreateTicketLogClient(c config) *CreateTicketLogClient {
	return &CreateTicketLogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `createticketlog.Hooks(f(g(h())))`.
func (c *CreateTicketLogClient) Use(hooks ...Hook) {
	c.hooks.CreateTicketLog = append(c.hooks.CreateTicketLog, hooks...)
}

// Create returns a builder for creating a CreateTicketLog entity.
func (c *CreateTicketLogClient) Create() *CreateTicketLogCreate {
	mutation := newCreateTicketLogMutation(c.config, OpCreate)
	return &CreateTicketLogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CreateTicketLog entities.
func (c *CreateTicketLogClient) CreateBulk(builders ...*CreateTicketLogCreate) *CreateTicketLogCreateBulk {
	return &CreateTicketLogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CreateTicketLog.
func (c *CreateTicketLogClient) Update() *CreateTicketLogUpdate {
	mutation := newCreateTicketLogMutation(c.config, OpUpdate)
	return &CreateTicketLogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CreateTicketLogClient) UpdateOne(ctl *CreateTicketLog) *CreateTicketLogUpdateOne {
	mutation := newCreateTicketLogMutation(c.config, OpUpdateOne, withCreateTicketLog(ctl))
	return &CreateTicketLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CreateTicketLogClient) UpdateOneID(id int) *CreateTicketLogUpdateOne {
	mutation := newCreateTicketLogMutation(c.config, OpUpdateOne, withCreateTicketLogID(id))
	return &CreateTicketLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CreateTicketLog.
func (c *CreateTicketLogClient) Delete() *CreateTicketLogDelete {
	mutation := newCreateTicketLogMutation(c.config, OpDelete)
	return &CreateTicketLogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CreateTicketLogClient) DeleteOne(ctl *CreateTicketLog) *CreateTicketLogDeleteOne {
	return c.DeleteOneID(ctl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CreateTicketLogClient) DeleteOneID(id int) *CreateTicketLogDeleteOne {
	builder := c.Delete().Where(createticketlog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CreateTicketLogDeleteOne{builder}
}

// Query returns a query builder for CreateTicketLog.
func (c *CreateTicketLogClient) Query() *CreateTicketLogQuery {
	return &CreateTicketLogQuery{
		config: c.config,
	}
}

// Get returns a CreateTicketLog entity by its id.
func (c *CreateTicketLogClient) Get(ctx context.Context, id int) (*CreateTicketLog, error) {
	return c.Query().Where(createticketlog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CreateTicketLogClient) GetX(ctx context.Context, id int) *CreateTicketLog {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CreateTicketLogClient) Hooks() []Hook {
	return c.hooks.CreateTicketLog
}

// ServiceProdiverClient is a client for the ServiceProdiver schema.
type ServiceProdiverClient struct {
	config
}

// NewServiceProdiverClient returns a client for the ServiceProdiver from the given config.
func NewServiceProdiverClient(c config) *ServiceProdiverClient {
	return &ServiceProdiverClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `serviceprodiver.Hooks(f(g(h())))`.
func (c *ServiceProdiverClient) Use(hooks ...Hook) {
	c.hooks.ServiceProdiver = append(c.hooks.ServiceProdiver, hooks...)
}

// Create returns a builder for creating a ServiceProdiver entity.
func (c *ServiceProdiverClient) Create() *ServiceProdiverCreate {
	mutation := newServiceProdiverMutation(c.config, OpCreate)
	return &ServiceProdiverCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ServiceProdiver entities.
func (c *ServiceProdiverClient) CreateBulk(builders ...*ServiceProdiverCreate) *ServiceProdiverCreateBulk {
	return &ServiceProdiverCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ServiceProdiver.
func (c *ServiceProdiverClient) Update() *ServiceProdiverUpdate {
	mutation := newServiceProdiverMutation(c.config, OpUpdate)
	return &ServiceProdiverUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServiceProdiverClient) UpdateOne(sp *ServiceProdiver) *ServiceProdiverUpdateOne {
	mutation := newServiceProdiverMutation(c.config, OpUpdateOne, withServiceProdiver(sp))
	return &ServiceProdiverUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServiceProdiverClient) UpdateOneID(id uuid.UUID) *ServiceProdiverUpdateOne {
	mutation := newServiceProdiverMutation(c.config, OpUpdateOne, withServiceProdiverID(id))
	return &ServiceProdiverUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ServiceProdiver.
func (c *ServiceProdiverClient) Delete() *ServiceProdiverDelete {
	mutation := newServiceProdiverMutation(c.config, OpDelete)
	return &ServiceProdiverDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ServiceProdiverClient) DeleteOne(sp *ServiceProdiver) *ServiceProdiverDeleteOne {
	return c.DeleteOneID(sp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ServiceProdiverClient) DeleteOneID(id uuid.UUID) *ServiceProdiverDeleteOne {
	builder := c.Delete().Where(serviceprodiver.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServiceProdiverDeleteOne{builder}
}

// Query returns a query builder for ServiceProdiver.
func (c *ServiceProdiverClient) Query() *ServiceProdiverQuery {
	return &ServiceProdiverQuery{
		config: c.config,
	}
}

// Get returns a ServiceProdiver entity by its id.
func (c *ServiceProdiverClient) Get(ctx context.Context, id uuid.UUID) (*ServiceProdiver, error) {
	return c.Query().Where(serviceprodiver.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServiceProdiverClient) GetX(ctx context.Context, id uuid.UUID) *ServiceProdiver {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ServiceProdiverClient) Hooks() []Hook {
	return c.hooks.ServiceProdiver
}

// TicketClient is a client for the Ticket schema.
type TicketClient struct {
	config
}

// NewTicketClient returns a client for the Ticket from the given config.
func NewTicketClient(c config) *TicketClient {
	return &TicketClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ticket.Hooks(f(g(h())))`.
func (c *TicketClient) Use(hooks ...Hook) {
	c.hooks.Ticket = append(c.hooks.Ticket, hooks...)
}

// Create returns a builder for creating a Ticket entity.
func (c *TicketClient) Create() *TicketCreate {
	mutation := newTicketMutation(c.config, OpCreate)
	return &TicketCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ticket entities.
func (c *TicketClient) CreateBulk(builders ...*TicketCreate) *TicketCreateBulk {
	return &TicketCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ticket.
func (c *TicketClient) Update() *TicketUpdate {
	mutation := newTicketMutation(c.config, OpUpdate)
	return &TicketUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TicketClient) UpdateOne(t *Ticket) *TicketUpdateOne {
	mutation := newTicketMutation(c.config, OpUpdateOne, withTicket(t))
	return &TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TicketClient) UpdateOneID(id uuid.UUID) *TicketUpdateOne {
	mutation := newTicketMutation(c.config, OpUpdateOne, withTicketID(id))
	return &TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ticket.
func (c *TicketClient) Delete() *TicketDelete {
	mutation := newTicketMutation(c.config, OpDelete)
	return &TicketDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TicketClient) DeleteOne(t *Ticket) *TicketDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TicketClient) DeleteOneID(id uuid.UUID) *TicketDeleteOne {
	builder := c.Delete().Where(ticket.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TicketDeleteOne{builder}
}

// Query returns a query builder for Ticket.
func (c *TicketClient) Query() *TicketQuery {
	return &TicketQuery{
		config: c.config,
	}
}

// Get returns a Ticket entity by its id.
func (c *TicketClient) Get(ctx context.Context, id uuid.UUID) (*Ticket, error) {
	return c.Query().Where(ticket.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TicketClient) GetX(ctx context.Context, id uuid.UUID) *Ticket {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Ticket.
func (c *TicketClient) QueryUser(t *Ticket) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticket.UserTable, ticket.UserColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLastEvent queries the last_event edge of a Ticket.
func (c *TicketClient) QueryLastEvent(t *Ticket) *TicketEventQuery {
	query := &TicketEventQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(ticketevent.Table, ticketevent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ticket.LastEventTable, ticket.LastEventColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTicketEvents queries the ticket_events edge of a Ticket.
func (c *TicketClient) QueryTicketEvents(t *Ticket) *TicketEventQuery {
	query := &TicketEventQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(ticketevent.Table, ticketevent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ticket.TicketEventsTable, ticket.TicketEventsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TicketClient) Hooks() []Hook {
	return c.hooks.Ticket
}

// TicketEventClient is a client for the TicketEvent schema.
type TicketEventClient struct {
	config
}

// NewTicketEventClient returns a client for the TicketEvent from the given config.
func NewTicketEventClient(c config) *TicketEventClient {
	return &TicketEventClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ticketevent.Hooks(f(g(h())))`.
func (c *TicketEventClient) Use(hooks ...Hook) {
	c.hooks.TicketEvent = append(c.hooks.TicketEvent, hooks...)
}

// Create returns a builder for creating a TicketEvent entity.
func (c *TicketEventClient) Create() *TicketEventCreate {
	mutation := newTicketEventMutation(c.config, OpCreate)
	return &TicketEventCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TicketEvent entities.
func (c *TicketEventClient) CreateBulk(builders ...*TicketEventCreate) *TicketEventCreateBulk {
	return &TicketEventCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TicketEvent.
func (c *TicketEventClient) Update() *TicketEventUpdate {
	mutation := newTicketEventMutation(c.config, OpUpdate)
	return &TicketEventUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TicketEventClient) UpdateOne(te *TicketEvent) *TicketEventUpdateOne {
	mutation := newTicketEventMutation(c.config, OpUpdateOne, withTicketEvent(te))
	return &TicketEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TicketEventClient) UpdateOneID(id uuid.UUID) *TicketEventUpdateOne {
	mutation := newTicketEventMutation(c.config, OpUpdateOne, withTicketEventID(id))
	return &TicketEventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TicketEvent.
func (c *TicketEventClient) Delete() *TicketEventDelete {
	mutation := newTicketEventMutation(c.config, OpDelete)
	return &TicketEventDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TicketEventClient) DeleteOne(te *TicketEvent) *TicketEventDeleteOne {
	return c.DeleteOneID(te.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TicketEventClient) DeleteOneID(id uuid.UUID) *TicketEventDeleteOne {
	builder := c.Delete().Where(ticketevent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TicketEventDeleteOne{builder}
}

// Query returns a query builder for TicketEvent.
func (c *TicketEventClient) Query() *TicketEventQuery {
	return &TicketEventQuery{
		config: c.config,
	}
}

// Get returns a TicketEvent entity by its id.
func (c *TicketEventClient) Get(ctx context.Context, id uuid.UUID) (*TicketEvent, error) {
	return c.Query().Where(ticketevent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TicketEventClient) GetX(ctx context.Context, id uuid.UUID) *TicketEvent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a TicketEvent.
func (c *TicketEventClient) QueryUser(te *TicketEvent) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := te.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticketevent.Table, ticketevent.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticketevent.UserTable, ticketevent.UserColumn),
		)
		fromV = sqlgraph.Neighbors(te.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTicket queries the ticket edge of a TicketEvent.
func (c *TicketEventClient) QueryTicket(te *TicketEvent) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := te.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticketevent.Table, ticketevent.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticketevent.TicketTable, ticketevent.TicketColumn),
		)
		fromV = sqlgraph.Neighbors(te.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TicketEventClient) Hooks() []Hook {
	return c.hooks.TicketEvent
}

// TransactionClient is a client for the Transaction schema.
type TransactionClient struct {
	config
}

// NewTransactionClient returns a client for the Transaction from the given config.
func NewTransactionClient(c config) *TransactionClient {
	return &TransactionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transaction.Hooks(f(g(h())))`.
func (c *TransactionClient) Use(hooks ...Hook) {
	c.hooks.Transaction = append(c.hooks.Transaction, hooks...)
}

// Create returns a builder for creating a Transaction entity.
func (c *TransactionClient) Create() *TransactionCreate {
	mutation := newTransactionMutation(c.config, OpCreate)
	return &TransactionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Transaction entities.
func (c *TransactionClient) CreateBulk(builders ...*TransactionCreate) *TransactionCreateBulk {
	return &TransactionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Transaction.
func (c *TransactionClient) Update() *TransactionUpdate {
	mutation := newTransactionMutation(c.config, OpUpdate)
	return &TransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransactionClient) UpdateOne(t *Transaction) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransaction(t))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransactionClient) UpdateOneID(id uuid.UUID) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransactionID(id))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Transaction.
func (c *TransactionClient) Delete() *TransactionDelete {
	mutation := newTransactionMutation(c.config, OpDelete)
	return &TransactionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TransactionClient) DeleteOne(t *Transaction) *TransactionDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TransactionClient) DeleteOneID(id uuid.UUID) *TransactionDeleteOne {
	builder := c.Delete().Where(transaction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransactionDeleteOne{builder}
}

// Query returns a query builder for Transaction.
func (c *TransactionClient) Query() *TransactionQuery {
	return &TransactionQuery{
		config: c.config,
	}
}

// Get returns a Transaction entity by its id.
func (c *TransactionClient) Get(ctx context.Context, id uuid.UUID) (*Transaction, error) {
	return c.Query().Where(transaction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransactionClient) GetX(ctx context.Context, id uuid.UUID) *Transaction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TransactionClient) Hooks() []Hook {
	return c.hooks.Transaction
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTickets queries the tickets edge of a User.
func (c *UserClient) QueryTickets(u *User) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.TicketsTable, user.TicketsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTicketEvents queries the ticket_events edge of a User.
func (c *UserClient) QueryTicketEvents(u *User) *TicketEventQuery {
	query := &TicketEventQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(ticketevent.Table, ticketevent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.TicketEventsTable, user.TicketEventsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
