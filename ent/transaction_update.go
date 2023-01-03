// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database-concurrency/ent/predicate"
	"database-concurrency/ent/transaction"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransactionUpdate is the builder for updating Transaction entities.
type TransactionUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionMutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tu *TransactionUpdate) Where(ps ...predicate.Transaction) *TransactionUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetHash sets the "hash" field.
func (tu *TransactionUpdate) SetHash(s string) *TransactionUpdate {
	tu.mutation.SetHash(s)
	return tu
}

// SetTime sets the "time" field.
func (tu *TransactionUpdate) SetTime(t time.Time) *TransactionUpdate {
	tu.mutation.SetTime(t)
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TransactionUpdate) SetCreatedAt(t time.Time) *TransactionUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableCreatedAt(t *time.Time) *TransactionUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// Mutation returns the TransactionMutation object of the builder.
func (tu *TransactionUpdate) Mutation() *TransactionMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transaction.Table,
			Columns: transaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transaction.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Hash(); ok {
		_spec.SetField(transaction.FieldHash, field.TypeString, value)
	}
	if value, ok := tu.mutation.Time(); ok {
		_spec.SetField(transaction.FieldTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(transaction.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TransactionUpdateOne is the builder for updating a single Transaction entity.
type TransactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionMutation
}

// SetHash sets the "hash" field.
func (tuo *TransactionUpdateOne) SetHash(s string) *TransactionUpdateOne {
	tuo.mutation.SetHash(s)
	return tuo
}

// SetTime sets the "time" field.
func (tuo *TransactionUpdateOne) SetTime(t time.Time) *TransactionUpdateOne {
	tuo.mutation.SetTime(t)
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TransactionUpdateOne) SetCreatedAt(t time.Time) *TransactionUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableCreatedAt(t *time.Time) *TransactionUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// Mutation returns the TransactionMutation object of the builder.
func (tuo *TransactionUpdateOne) Mutation() *TransactionMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TransactionUpdateOne) Select(field string, fields ...string) *TransactionUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Transaction entity.
func (tuo *TransactionUpdateOne) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionUpdateOne) SaveX(ctx context.Context) *Transaction {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TransactionUpdateOne) sqlSave(ctx context.Context) (_node *Transaction, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transaction.Table,
			Columns: transaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transaction.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Transaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.FieldID)
		for _, f := range fields {
			if !transaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Hash(); ok {
		_spec.SetField(transaction.FieldHash, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Time(); ok {
		_spec.SetField(transaction.FieldTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(transaction.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Transaction{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}