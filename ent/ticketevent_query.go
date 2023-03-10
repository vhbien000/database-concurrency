// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database-concurrency/ent/predicate"
	"database-concurrency/ent/ticket"
	"database-concurrency/ent/ticketevent"
	"database-concurrency/ent/user"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TicketEventQuery is the builder for querying TicketEvent entities.
type TicketEventQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TicketEvent
	withUser   *UserQuery
	withTicket *TicketQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TicketEventQuery builder.
func (teq *TicketEventQuery) Where(ps ...predicate.TicketEvent) *TicketEventQuery {
	teq.predicates = append(teq.predicates, ps...)
	return teq
}

// Limit adds a limit step to the query.
func (teq *TicketEventQuery) Limit(limit int) *TicketEventQuery {
	teq.limit = &limit
	return teq
}

// Offset adds an offset step to the query.
func (teq *TicketEventQuery) Offset(offset int) *TicketEventQuery {
	teq.offset = &offset
	return teq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (teq *TicketEventQuery) Unique(unique bool) *TicketEventQuery {
	teq.unique = &unique
	return teq
}

// Order adds an order step to the query.
func (teq *TicketEventQuery) Order(o ...OrderFunc) *TicketEventQuery {
	teq.order = append(teq.order, o...)
	return teq
}

// QueryUser chains the current query on the "user" edge.
func (teq *TicketEventQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: teq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := teq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := teq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ticketevent.Table, ticketevent.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticketevent.UserTable, ticketevent.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(teq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTicket chains the current query on the "ticket" edge.
func (teq *TicketEventQuery) QueryTicket() *TicketQuery {
	query := &TicketQuery{config: teq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := teq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := teq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ticketevent.Table, ticketevent.FieldID, selector),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticketevent.TicketTable, ticketevent.TicketColumn),
		)
		fromU = sqlgraph.SetNeighbors(teq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TicketEvent entity from the query.
// Returns a *NotFoundError when no TicketEvent was found.
func (teq *TicketEventQuery) First(ctx context.Context) (*TicketEvent, error) {
	nodes, err := teq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ticketevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (teq *TicketEventQuery) FirstX(ctx context.Context) *TicketEvent {
	node, err := teq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TicketEvent ID from the query.
// Returns a *NotFoundError when no TicketEvent ID was found.
func (teq *TicketEventQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = teq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ticketevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (teq *TicketEventQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := teq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TicketEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TicketEvent entity is found.
// Returns a *NotFoundError when no TicketEvent entities are found.
func (teq *TicketEventQuery) Only(ctx context.Context) (*TicketEvent, error) {
	nodes, err := teq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ticketevent.Label}
	default:
		return nil, &NotSingularError{ticketevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (teq *TicketEventQuery) OnlyX(ctx context.Context) *TicketEvent {
	node, err := teq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TicketEvent ID in the query.
// Returns a *NotSingularError when more than one TicketEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (teq *TicketEventQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = teq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ticketevent.Label}
	default:
		err = &NotSingularError{ticketevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (teq *TicketEventQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := teq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TicketEvents.
func (teq *TicketEventQuery) All(ctx context.Context) ([]*TicketEvent, error) {
	if err := teq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return teq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (teq *TicketEventQuery) AllX(ctx context.Context) []*TicketEvent {
	nodes, err := teq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TicketEvent IDs.
func (teq *TicketEventQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := teq.Select(ticketevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (teq *TicketEventQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := teq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (teq *TicketEventQuery) Count(ctx context.Context) (int, error) {
	if err := teq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return teq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (teq *TicketEventQuery) CountX(ctx context.Context) int {
	count, err := teq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (teq *TicketEventQuery) Exist(ctx context.Context) (bool, error) {
	if err := teq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return teq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (teq *TicketEventQuery) ExistX(ctx context.Context) bool {
	exist, err := teq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TicketEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (teq *TicketEventQuery) Clone() *TicketEventQuery {
	if teq == nil {
		return nil
	}
	return &TicketEventQuery{
		config:     teq.config,
		limit:      teq.limit,
		offset:     teq.offset,
		order:      append([]OrderFunc{}, teq.order...),
		predicates: append([]predicate.TicketEvent{}, teq.predicates...),
		withUser:   teq.withUser.Clone(),
		withTicket: teq.withTicket.Clone(),
		// clone intermediate query.
		sql:    teq.sql.Clone(),
		path:   teq.path,
		unique: teq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (teq *TicketEventQuery) WithUser(opts ...func(*UserQuery)) *TicketEventQuery {
	query := &UserQuery{config: teq.config}
	for _, opt := range opts {
		opt(query)
	}
	teq.withUser = query
	return teq
}

// WithTicket tells the query-builder to eager-load the nodes that are connected to
// the "ticket" edge. The optional arguments are used to configure the query builder of the edge.
func (teq *TicketEventQuery) WithTicket(opts ...func(*TicketQuery)) *TicketEventQuery {
	query := &TicketQuery{config: teq.config}
	for _, opt := range opts {
		opt(query)
	}
	teq.withTicket = query
	return teq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TicketID uuid.UUID `json:"ticket_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TicketEvent.Query().
//		GroupBy(ticketevent.FieldTicketID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (teq *TicketEventQuery) GroupBy(field string, fields ...string) *TicketEventGroupBy {
	grbuild := &TicketEventGroupBy{config: teq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := teq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return teq.sqlQuery(ctx), nil
	}
	grbuild.label = ticketevent.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TicketID uuid.UUID `json:"ticket_id,omitempty"`
//	}
//
//	client.TicketEvent.Query().
//		Select(ticketevent.FieldTicketID).
//		Scan(ctx, &v)
func (teq *TicketEventQuery) Select(fields ...string) *TicketEventSelect {
	teq.fields = append(teq.fields, fields...)
	selbuild := &TicketEventSelect{TicketEventQuery: teq}
	selbuild.label = ticketevent.Label
	selbuild.flds, selbuild.scan = &teq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a TicketEventSelect configured with the given aggregations.
func (teq *TicketEventQuery) Aggregate(fns ...AggregateFunc) *TicketEventSelect {
	return teq.Select().Aggregate(fns...)
}

func (teq *TicketEventQuery) prepareQuery(ctx context.Context) error {
	for _, f := range teq.fields {
		if !ticketevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if teq.path != nil {
		prev, err := teq.path(ctx)
		if err != nil {
			return err
		}
		teq.sql = prev
	}
	return nil
}

func (teq *TicketEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TicketEvent, error) {
	var (
		nodes       = []*TicketEvent{}
		_spec       = teq.querySpec()
		loadedTypes = [2]bool{
			teq.withUser != nil,
			teq.withTicket != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TicketEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TicketEvent{config: teq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(teq.modifiers) > 0 {
		_spec.Modifiers = teq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, teq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := teq.withUser; query != nil {
		if err := teq.loadUser(ctx, query, nodes, nil,
			func(n *TicketEvent, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := teq.withTicket; query != nil {
		if err := teq.loadTicket(ctx, query, nodes, nil,
			func(n *TicketEvent, e *Ticket) { n.Edges.Ticket = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (teq *TicketEventQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*TicketEvent, init func(*TicketEvent), assign func(*TicketEvent, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*TicketEvent)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (teq *TicketEventQuery) loadTicket(ctx context.Context, query *TicketQuery, nodes []*TicketEvent, init func(*TicketEvent), assign func(*TicketEvent, *Ticket)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*TicketEvent)
	for i := range nodes {
		fk := nodes[i].TicketID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(ticket.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ticket_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (teq *TicketEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := teq.querySpec()
	if len(teq.modifiers) > 0 {
		_spec.Modifiers = teq.modifiers
	}
	_spec.Node.Columns = teq.fields
	if len(teq.fields) > 0 {
		_spec.Unique = teq.unique != nil && *teq.unique
	}
	return sqlgraph.CountNodes(ctx, teq.driver, _spec)
}

func (teq *TicketEventQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := teq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (teq *TicketEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ticketevent.Table,
			Columns: ticketevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ticketevent.FieldID,
			},
		},
		From:   teq.sql,
		Unique: true,
	}
	if unique := teq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := teq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ticketevent.FieldID)
		for i := range fields {
			if fields[i] != ticketevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := teq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := teq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := teq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := teq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (teq *TicketEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(teq.driver.Dialect())
	t1 := builder.Table(ticketevent.Table)
	columns := teq.fields
	if len(columns) == 0 {
		columns = ticketevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if teq.sql != nil {
		selector = teq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if teq.unique != nil && *teq.unique {
		selector.Distinct()
	}
	for _, m := range teq.modifiers {
		m(selector)
	}
	for _, p := range teq.predicates {
		p(selector)
	}
	for _, p := range teq.order {
		p(selector)
	}
	if offset := teq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := teq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (teq *TicketEventQuery) ForUpdate(opts ...sql.LockOption) *TicketEventQuery {
	if teq.driver.Dialect() == dialect.Postgres {
		teq.Unique(false)
	}
	teq.modifiers = append(teq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return teq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (teq *TicketEventQuery) ForShare(opts ...sql.LockOption) *TicketEventQuery {
	if teq.driver.Dialect() == dialect.Postgres {
		teq.Unique(false)
	}
	teq.modifiers = append(teq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return teq
}

// TicketEventGroupBy is the group-by builder for TicketEvent entities.
type TicketEventGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tegb *TicketEventGroupBy) Aggregate(fns ...AggregateFunc) *TicketEventGroupBy {
	tegb.fns = append(tegb.fns, fns...)
	return tegb
}

// Scan applies the group-by query and scans the result into the given value.
func (tegb *TicketEventGroupBy) Scan(ctx context.Context, v any) error {
	query, err := tegb.path(ctx)
	if err != nil {
		return err
	}
	tegb.sql = query
	return tegb.sqlScan(ctx, v)
}

func (tegb *TicketEventGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range tegb.fields {
		if !ticketevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tegb *TicketEventGroupBy) sqlQuery() *sql.Selector {
	selector := tegb.sql.Select()
	aggregation := make([]string, 0, len(tegb.fns))
	for _, fn := range tegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tegb.fields)+len(tegb.fns))
		for _, f := range tegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tegb.fields...)...)
}

// TicketEventSelect is the builder for selecting fields of TicketEvent entities.
type TicketEventSelect struct {
	*TicketEventQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tes *TicketEventSelect) Aggregate(fns ...AggregateFunc) *TicketEventSelect {
	tes.fns = append(tes.fns, fns...)
	return tes
}

// Scan applies the selector query and scans the result into the given value.
func (tes *TicketEventSelect) Scan(ctx context.Context, v any) error {
	if err := tes.prepareQuery(ctx); err != nil {
		return err
	}
	tes.sql = tes.TicketEventQuery.sqlQuery(ctx)
	return tes.sqlScan(ctx, v)
}

func (tes *TicketEventSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(tes.fns))
	for _, fn := range tes.fns {
		aggregation = append(aggregation, fn(tes.sql))
	}
	switch n := len(*tes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		tes.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		tes.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := tes.sql.Query()
	if err := tes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
