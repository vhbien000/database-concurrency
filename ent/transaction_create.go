// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database-concurrency/ent/transaction"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
}

// SetHash sets the "hash" field.
func (tc *TransactionCreate) SetHash(s string) *TransactionCreate {
	tc.mutation.SetHash(s)
	return tc
}

// SetTime sets the "time" field.
func (tc *TransactionCreate) SetTime(t time.Time) *TransactionCreate {
	tc.mutation.SetTime(t)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TransactionCreate) SetCreatedAt(t time.Time) *TransactionCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableCreatedAt(t *time.Time) *TransactionCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionCreate) SetID(u uuid.UUID) *TransactionCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableID(u *uuid.UUID) *TransactionCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
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
		nv, ok := v.(*Transaction)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TransactionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TransactionCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := transaction.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := transaction.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Transaction.hash"`)}
	}
	if _, ok := tc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "Transaction.time"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Transaction.created_at"`)}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
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

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transaction.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.Hash(); ok {
		_spec.SetField(transaction.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if value, ok := tc.mutation.Time(); ok {
		_spec.SetField(transaction.FieldTime, field.TypeTime, value)
		_node.Time = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(transaction.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	builders []*TransactionCreate
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
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
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
