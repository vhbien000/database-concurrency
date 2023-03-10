// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database-concurrency/ent/ticket"
	"database-concurrency/ent/ticketevent"
	"database-concurrency/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TicketCreate is the builder for creating a Ticket entity.
type TicketCreate struct {
	config
	mutation *TicketMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (tc *TicketCreate) SetStatus(s string) *TicketCreate {
	tc.mutation.SetStatus(s)
	return tc
}

// SetUserID sets the "user_id" field.
func (tc *TicketCreate) SetUserID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetUserID(u)
	return tc
}

// SetMetadata sets the "metadata" field.
func (tc *TicketCreate) SetMetadata(m map[string]interface{}) *TicketCreate {
	tc.mutation.SetMetadata(m)
	return tc
}

// SetVersions sets the "versions" field.
func (tc *TicketCreate) SetVersions(s string) *TicketCreate {
	tc.mutation.SetVersions(s)
	return tc
}

// SetNillableVersions sets the "versions" field if the given value is not nil.
func (tc *TicketCreate) SetNillableVersions(s *string) *TicketCreate {
	if s != nil {
		tc.SetVersions(*s)
	}
	return tc
}

// SetLastEventID sets the "last_event_id" field.
func (tc *TicketCreate) SetLastEventID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetLastEventID(u)
	return tc
}

// SetNillableLastEventID sets the "last_event_id" field if the given value is not nil.
func (tc *TicketCreate) SetNillableLastEventID(u *uuid.UUID) *TicketCreate {
	if u != nil {
		tc.SetLastEventID(*u)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TicketCreate) SetCreatedAt(t time.Time) *TicketCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TicketCreate) SetNillableCreatedAt(t *time.Time) *TicketCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TicketCreate) SetUpdatedAt(t time.Time) *TicketCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TicketCreate) SetNillableUpdatedAt(t *time.Time) *TicketCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TicketCreate) SetID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TicketCreate) SetNillableID(u *uuid.UUID) *TicketCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// SetUser sets the "user" edge to the User entity.
func (tc *TicketCreate) SetUser(u *User) *TicketCreate {
	return tc.SetUserID(u.ID)
}

// SetLastEvent sets the "last_event" edge to the TicketEvent entity.
func (tc *TicketCreate) SetLastEvent(t *TicketEvent) *TicketCreate {
	return tc.SetLastEventID(t.ID)
}

// AddTicketEventIDs adds the "ticket_events" edge to the TicketEvent entity by IDs.
func (tc *TicketCreate) AddTicketEventIDs(ids ...uuid.UUID) *TicketCreate {
	tc.mutation.AddTicketEventIDs(ids...)
	return tc
}

// AddTicketEvents adds the "ticket_events" edges to the TicketEvent entity.
func (tc *TicketCreate) AddTicketEvents(t ...*TicketEvent) *TicketCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTicketEventIDs(ids...)
}

// Mutation returns the TicketMutation object of the builder.
func (tc *TicketCreate) Mutation() *TicketMutation {
	return tc.mutation
}

// Save creates the Ticket in the database.
func (tc *TicketCreate) Save(ctx context.Context) (*Ticket, error) {
	var (
		err  error
		node *Ticket
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TicketMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Ticket)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TicketMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TicketCreate) SaveX(ctx context.Context) *Ticket {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TicketCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TicketCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TicketCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := ticket.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := ticket.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := ticket.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TicketCreate) check() error {
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Ticket.status"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Ticket.user_id"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Ticket.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Ticket.updated_at"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Ticket.user"`)}
	}
	return nil
}

func (tc *TicketCreate) sqlSave(ctx context.Context) (*Ticket, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (tc *TicketCreate) createSpec() (*Ticket, *sqlgraph.CreateSpec) {
	var (
		_node = &Ticket{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ticket.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ticket.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(ticket.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.Metadata(); ok {
		_spec.SetField(ticket.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if value, ok := tc.mutation.Versions(); ok {
		_spec.SetField(ticket.FieldVersions, field.TypeString, value)
		_node.Versions = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(ticket.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(ticket.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.UserTable,
			Columns: []string{ticket.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.LastEventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   ticket.LastEventTable,
			Columns: []string{ticket.LastEventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticketevent.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.LastEventID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TicketEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ticket.TicketEventsTable,
			Columns: []string{ticket.TicketEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ticketevent.FieldID,
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

// TicketCreateBulk is the builder for creating many Ticket entities in bulk.
type TicketCreateBulk struct {
	config
	builders []*TicketCreate
}

// Save creates the Ticket entities in the database.
func (tcb *TicketCreateBulk) Save(ctx context.Context) ([]*Ticket, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Ticket, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TicketMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TicketCreateBulk) SaveX(ctx context.Context) []*Ticket {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TicketCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TicketCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
